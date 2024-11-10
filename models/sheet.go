package models

import "time"

type Sheet struct {
	SheetID     string    `json:"sheet_id"`
	Name        string    `json:"name"`
	Brand       string    `json:"coffee_brand"`
	MenuURL     string    `json:"menu_url"`
	HostIDs     string    `json:"host_user_id"`
	IsActive    bool      `json:"is_active"`
	CreateDatim time.Time `json:"created_at, omitempty"`
	Orders      Orders    `json:"orders, omitempty"`
}

type GetSheetReq struct {
	SheetID string `json:"sheet_id"`
}

func (req *GetSheetReq) GetSheetID() string {
	if req != nil {
		return req.SheetID
	}
	return "undefined"
}
