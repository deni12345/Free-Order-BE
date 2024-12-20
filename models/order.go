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

type GetSheetOrdersReq struct {
	SheetID string `json:"sheet_id"`
}

func (req *GetSheetOrdersReq) GetSheetID() string {
	if req != nil {
		return req.SheetID
	}
	return ""
}

type GetUserOrdersReq struct {
	UserID string `json:"user_id"`
}

func (req *GetUserOrdersReq) GetUserID() string {
	if req != nil {
		return req.UserID
	}
	return ""
}
