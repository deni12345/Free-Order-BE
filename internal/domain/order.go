package domain

import (
	"github/free-order-be/models"
	"time"
)

type Orders []*Order

type Order struct {
	PK          string    `dynamo:"PK,hash"`
	SK          string    `dynamo:"PK,hash"`
	Name        string    `dynamo:"Name"`
	UserID      string    `dynamo:"UserID"`
	Amount      uint      `dynamo:"Amount"`
	Price       uint      `dynamo:"Price"`
	IsActive    bool      `dynamo:"IsActive"`
	CreateDatim time.Time `dynamo:"CreateDatim"`
}

func (o *Order) GetPK() string {
	if o != nil {
		return o.PK
	}
	return ""
}

func (o *Order) GetSK() string {
	if o != nil {
		return o.SK
	}
	return ""
}

func (o *Order) CheckNil() *Order {
	if o.PK != "" {
		return o
	}
	return nil
}

func (o *Order) GetName() string {
	if o != nil {
		return o.Name
	}
	return ""
}

func (o *Order) GetUserID() string {
	if o != nil {
		return o.UserID
	}
	return ""
}

func (o *Order) GetAmount() uint {
	if o != nil {
		return o.Amount
	}
	return 0
}

func (o *Order) GetIsActive() bool {
	if o != nil {
		return o.IsActive
	}
	return false
}

func (o *Order) GetPrice() uint {
	if o != nil {
		return o.Price
	}
	return 0
}

func (o *Order) GetCreateDatim() time.Time {
	if o != nil {
		return o.CreateDatim
	}
	return time.Time{}
}

func (o *Order) GetModelOrder() *models.Order {
	return &models.Order{
		PK:          o.GetPK(),
		SK:          o.GetSK(),
		Name:        o.GetName(),
		UserID:      o.GetUserID(),
		Amount:      o.GetAmount(),
		Price:       o.GetPrice(),
		IsActive:    o.GetIsActive(),
		CreateDatim: o.GetCreateDatim(),
	}
}

func BuildDomainOrder(v *models.Order) *Order {
	if v == nil {
		return nil
	}

	return &Order{
		PK:          v.PK,
		SK:          v.SK,
		Name:        v.Name,
		UserID:      v.UserID,
		Amount:      v.Amount,
		Price:       v.Price,
		IsActive:    v.IsActive,
		CreateDatim: time.Now().UTC(),
	}
}
