package main

import (
	"encoding/json"
	"net/http"
)

func PlacementHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Request body
	var req struct {
		Yard            string  `json:"yard"`
		ContainerNumber string  `json:"container_number"`
		Block           string  `json:"block"`
		Slot            int     `json:"slot"`
		Row             int     `json:"row"`
		Tier            int     `json:"tier"`
		ContainerSize   int     `json:"container_size"`   // 20 atau 40
		ContainerHeight float64 `json:"container_height"` // 8.6 atau 9.6
		ContainerType   string  `json:"container_type"`   // DRY, REEFER, dll
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// -------------------------
	// 1. Validasi yard
	// -------------------------
	var yardID int64
	if err := db.QueryRow(`SELECT id FROM yards WHERE code=$1`, req.Yard).Scan(&yardID); err != nil {
		http.Error(w, "yard not found", http.StatusNotFound)
		return
	}

	// -------------------------
	// 2. Validasi block & ambil jumlah slot block
	// -------------------------
	var blockID, blockSlots int
	if err := db.QueryRow(`
		SELECT id, slots 
		FROM blocks 
		WHERE block_code=$1 AND yard_id=$2
	`, req.Block, yardID).Scan(&blockID, &blockSlots); err != nil {
		http.Error(w, "block not found", http.StatusNotFound)
		return
	}

	// -------------------------
	// 3. Tentukan slotNeededx
	// -------------------------
	slotNeeded := 1
	if req.ContainerSize == 40 {
		slotNeeded = 2
	} else if req.ContainerSize != 20 {
		http.Error(w, "invalid container_size, must be 20 or 40", http.StatusBadRequest)
		return
	}

	// -------------------------
	// 4. Cek apakah slot, row, tier valid
	// -------------------------
	if req.Slot < 1 || req.Slot+slotNeeded-1 > blockSlots {
		http.Error(w, "not enough slots in block", http.StatusBadRequest)
		return
	}
	if req.Row < 1 || req.Tier < 1 {
		http.Error(w, "invalid row or tier", http.StatusBadRequest)
		return
	}

	// -------------------------
	// 5. Cek apakah slot(s) kosong
	// -------------------------
	for s := 0; s < slotNeeded; s++ {
		var count int
		err := db.QueryRow(`
			SELECT COUNT(*) 
			FROM yard_plans 
			WHERE block_id=$1 AND slot=$2 AND row=$3 AND tier=$4
		`, blockID, req.Slot+s, req.Row, req.Tier).Scan(&count)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if count > 0 {
			http.Error(w, "slot(s) already occupied", http.StatusConflict)
			return
		}
	}

	// -------------------------
	// 6. Insert ke yard_plans
	// -------------------------
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for s := 0; s < slotNeeded; s++ {
		_, err := tx.Exec(`
			INSERT INTO yard_plans
			(block_id, slot, row, tier, container_number, container_size, container_height, container_type)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		`, blockID, req.Slot+s, req.Row, req.Tier, req.ContainerNumber, req.ContainerSize, req.ContainerHeight, req.ContainerType)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{
		"message": "container placed successfully",
	})
}
