package models

import "time"

type Sheet struct {
	PK          string    `json:"id"`
	Name        string    `json:"name"`
	Brand       string    `json:"coffee_brand"`
	MenuURL     string    `json:"menu_url"`
	HostIDs     string    `json:"host_user_id"`
	IsActive    bool      `json:"is_active"`
	CreateDatim time.Time `json:"created_at,omitempty"`
}
