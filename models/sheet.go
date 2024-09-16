package models

type Sheet struct {
	Name        string `json:"name"`
	CreatedBy   string `json:"created_by"`
	EndDate     string `json:"end_at"`
	CreatedDate string `json:"created_at,omitempty"`
	Orders      Orders
}

type Orders []*Order

type Order struct {
	OrderBy  string `json:"order_by"`
	FoodName string `json:"food_name"`
	Amount   int32  `json:"amount"`
	CreateAt string `json:"create_at"`
}
