package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func SuggestionHandler(w http.ResponseWriter, r *http.Request) {
	var req SuggestionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	// 1. CARI YARD
	var yardID int64
	err := db.QueryRow(`SELECT id FROM yards WHERE code = $1`, req.Yard).Scan(&yardID)
	if err == sql.ErrNoRows {
		writeJSON(w, map[string]string{"message": "yard not found"})
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 2. GET ALL BLOCKS IN THE YARD
	rows, err := db.Query(`
		SELECT id, block_code, slots, rows, tiers
		FROM blocks
		WHERE yard_id = $1
		ORDER BY block_code ASC
	`, yardID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	type blockStruct struct {
		ID        int
		BlockCode string
		Slots     int
		Rows      int
		Tiers     int
	}

	var blocks []blockStruct
	for rows.Next() {
		var b blockStruct
		rows.Scan(&b.ID, &b.BlockCode, &b.Slots, &b.Rows, &b.Tiers)
		blocks = append(blocks, b)
	}

	// Tentukan jumlah slot yang dibutuhkan
	slotNeeded := 1
	if req.ContainerSize == 40 {
		slotNeeded = 2
	}

	// 3. LOOP CARI POSISI KOSONG
	for _, b := range blocks {
		for tier := 1; tier <= b.Tiers; tier++ { // tier bawah → atas
			for slot := 1; slot <= b.Slots-slotNeeded+1; slot++ { // slot kiri → kanan
				for row := 1; row <= b.Rows; row++ { // row bawah → atas

					occupied := false
					for s := 0; s < slotNeeded; s++ {
						var count int
						err := db.QueryRow(`
							SELECT COUNT(*)
							FROM yard_plans
							WHERE block_id=$1 AND slot=$2 AND row=$3 AND tier=$4
							AND is_picked=FALSE
						`, b.ID, slot+s, row, tier).Scan(&count)

						if err != nil || count > 0 {
							occupied = true
							break
						}
					}

					if !occupied {
						writeJSON(w, map[string]interface{}{
							"suggested_position": map[string]interface{}{
								"block": b.BlockCode,
								"slot":  slot,
								"row":   row,
								"tier":  tier,
							},
						})
						return
					}
				}
			}
		}
	}

	writeJSON(w, map[string]string{
		"message": "no available slot",
	})
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
