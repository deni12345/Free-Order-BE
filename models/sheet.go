package models

import (
	"github/lambda-microservice/internal/domain"
	"time"
)

type Sheet struct {
	ID        *uint      `json:"id"`
	UserID    *uint      ``
	Name      string     `json:"name"`
	EndDate   string     `json:"end_at"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	IsActive  bool       `json:"is_active"`
	Orders    Orders
}

type Orders []*Order

type Order struct {
	ID       *uint  `json:"id"`
	OrderBy  string `json:"order_by"`
	FoodName string `json:"food_name"`
	Amount   int32  `json:"amount"`
	Note     string `json:"note"`
	CreateAt string `json:"create_at"`
}

func (sh *Sheet) GetID() *uint {
	if sh != nil {
		return sh.ID
	}
	return nil
}

func (sh *Sheet) GetName() string {
	if sh != nil {
		return sh.Name
	}
	return ""
}

func (sh *Sheet) GetCreatedBy() string {
	if sh != nil {
		return sh.Name
	}
	return ""
}

func (sh *Sheet) GetCreatedAt() *time.Time {
	if sh != nil {
		return sh.CreatedAt
	}
	return nil
}

func (sh *Sheet) BuildDomainUser() *domain.Sheet {
	if sh == nil {
		return nil
	}
	return &domain.Sheet{
		ID:       sh.GetID(),
		Name:     sh.GetName(),
		CreateAt: sh.GetCreatedAt(),
		UserId:   new(uint),
		OrderId:  new(uint),
		Order:    &domain.Order{},
	}
}
