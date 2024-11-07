package models

import "time"

type Orders []*Order

type Order struct {
	OrderID     string    `json:"order_id"`
	SheetID     string    `json:"sheet_id"`
	Name        string    `json:"name"`
	UserID      string    `json:"user_id"`
	Amount      uint      `json:"amount"`
	Price       uint      `json:"price"`
	IsActive    bool      `json:"is_active"`
	CreateDatim time.Time `json:"created_at,omitempty"`
}

type GetOrdersReq struct {
	SheetID string `json:"sheet_id"`
}

func (req *GetOrdersReq) GetSheetID() string {
	if req != nil {
		return req.SheetID
	}
	return ""
}
