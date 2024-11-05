package models

import "time"

type Order struct {
	SK          string
	PK          string
	Name        string
	UserID      string
	Amount      uint
	Price       uint
	IsActive    bool
	CreateDatim time.Time
}
