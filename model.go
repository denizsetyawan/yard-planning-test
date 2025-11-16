package main

import "time"

// -----------------------------
// Request / Response models
// -----------------------------

// SuggestionRequest adalah payload untuk /suggestion
type SuggestionRequest struct {
	Yard            string  `json:"yard"`
	ContainerNumber string  `json:"container_number"`
	ContainerSize   int     `json:"container_size"`
	ContainerHeight float64 `json:"container_height"`
	ContainerType   string  `json:"container_type"`
}

// SuggestedPosition adalah posisi yang direkomendasikan
type SuggestedPosition struct {
	Block string `json:"block"`
	Slot  int    `json:"slot"`
	Row   int    `json:"row"`
	Tier  int    `json:"tier"`
}

// SuggestionResponse response untuk /suggestion
type SuggestionResponse struct {
	SuggestedPosition *SuggestedPosition `json:"suggested_position,omitempty"`
	Message           string             `json:"message,omitempty"`
}

// -----------------------------
// Database entity models
// -----------------------------

// Yard merepresentasikan tabel yards
type Yard struct {
	ID        int64     `json:"id"`          // bigint/int8
	Code      string    `json:"code"`        // varchar
	Name      string    `json:"name"`        // varchar
	Desc      string    `json:"description"` // text
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Block merepresentasikan tabel blocks
type Block struct {
	ID        int       `json:"id"`         // int4
	YardID    int       `json:"yard_id"`    // int4
	BlockCode string    `json:"block_code"` // varchar
	Slots     int       `json:"slots"`      // int4
	Rows      int       `json:"rows"`       // int4
	Tiers     int       `json:"tiers"`      // int4
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// YardPlan merepresentasikan tabel yard_plans (satu record = satu posisi / cell)
type YardPlan struct {
	ID              int       `json:"id"`               // int4
	BlockID         int       `json:"block_id"`         // int4 (FK ke blocks.id)
	Slot            int       `json:"slot"`             // int4
	Row             int       `json:"row"`              // int4
	Tier            int       `json:"tier"`             // int4
	ContainerNumber string    `json:"container_number"` // varchar (nullable)
	ContainerSize   int       `json:"container_size"`   // int4 (nullable)
	ContainerHeight float64   `json:"container_height"` // numeric (nullable)
	ContainerType   string    `json:"container_type"`   // varchar (nullable)
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
