package models

type Sheet struct {
	ID          *uint    `json:"id"`
	Name        string   `json:"name"`
	CoffeeBrand string   `json:"coffee_brand"`
	MenuURL     string   `json:"menu_url"`
	CoopHost    []string `json:"coop_host"`
	HostUserID  uint     `json:"host_user_id"`
	IsActive    bool     `json:"is_active"`
}
