package domain

import (
	"github/free-order-be/models"
	"time"
)

type Orders []*Order

type Order struct {
	PK       string    `dynamo:"PK,hash"`
	SK       string    `dynamo:"SK,hash"`
	Name     string    `dynamo:"Name"`
	UserID   string    `dynamo:"UserID"`
	Amount   uint      `dynamo:"Amount"`
	Price    uint      `dynamo:"Price"`
	IsActive bool      `dynamo:"IsActive"`
	CreateAt time.Time `dynamo:"CreateAt"`
}

func (o *Order) GetPK() string {
	if o != nil && o.PK != "" {
		return o.PK
	}
	return UNDEFINED
}

func (o *Order) GetSK() string {
	if o != nil && o.SK != "" {
		return o.SK
	}
	return UNDEFINED
}

func (o *Order) IsValid() *Order {
	if o.GetPK() != UNDEFINED {
		return o
	}
	return nil
}

func (o *Order) IsNil() bool {
	if o.GetSK() != UNDEFINED {
		return false
	}
	return true
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

func (o *Order) GetCreateAt() time.Time {
	if o != nil {
		return o.CreateAt
	}
	return time.Time{}
}

func (o *Order) GetModelOrder() *models.Order {
	return &models.Order{
		SheetID:     o.GetPK(),
		OrderID:     o.GetSK(),
		Name:        o.GetName(),
		UserID:      o.GetUserID(),
		Amount:      o.GetAmount(),
		Price:       o.GetPrice(),
		IsActive:    o.GetIsActive(),
		CreateDatim: o.GetCreateAt(),
	}
}

func (os Orders) GetModelOrders() models.Orders {
	if len(os) == 0 {
		return nil
	}

	orders := make(models.Orders, 0, len(os))
	for _, o := range os {
		orders = append(orders, o.GetModelOrder())
	}
	return orders
}

func BuildDomainOrder(v *models.Order) *Order {
	if v == nil {
		return nil
	}

	return &Order{
		PK:       v.SheetID,
		SK:       v.OrderID,
		Name:     v.Name,
		UserID:   v.UserID,
		Amount:   v.Amount,
		Price:    v.Price,
		IsActive: v.IsActive,
		CreateAt: time.Now().UTC(),
	}
}
