package models

import "time"

type Order struct {
	SK          string    `json:"id"`
	PK          string    `json:"sheet_id"`
	Name        string    `json:"name"`
	UserID      string    `json:"user_id"`
	Amount      uint      `json:"amount"`
	Price       uint      `json:"price"`
	IsActive    bool      `json:"is_active"`
	CreateDatim time.Time `json:"created_at,omitempty"`
}
