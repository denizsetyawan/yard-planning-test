package main

import (
	"encoding/json"
	"net/http"
)

func PickupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Yard            string `json:"yard"`
		ContainerNumber string `json:"container_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// 1. Cari yard
	var yardID int64
	if err := db.QueryRow(`SELECT id FROM yards WHERE code=$1`, req.Yard).Scan(&yardID); err != nil {
		http.Error(w, "yard not found", http.StatusNotFound)
		return
	}

	// 2. Update is_picked = true
	res, err := db.Exec(`
		UPDATE yard_plans
		SET is_picked = TRUE
		WHERE container_number = $1
		AND block_id IN (
			SELECT id FROM blocks WHERE yard_id = (
				SELECT id FROM yards WHERE code = $2
			)
		)
		AND is_picked = FALSE
	`, req.ContainerNumber, req.Yard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		http.Error(w, "container not found or already picked up", http.StatusNotFound)
		return
	}

	writeJSON(w, map[string]string{
		"message": "container picked up successfully",
	})
}
