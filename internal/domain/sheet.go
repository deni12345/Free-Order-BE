package domain

import (
	"github/free-order-be/models"
	"time"
)

const (
	SheetTable = "Sheet"
)

type Sheets []*Sheet

type Sheet struct {
	PK       *string   `dynamo:"PK,hash"`
	SK       *string   `dynamo:"SK,range"`
	Name     string    `dynamo:"Name"`
	Brand    string    `dynamo:"Brand"`
	MenuURL  string    `dynamo:"MenuURL"`
	HostIDs  string    `dynamo:"HostUserId"`
	IsActive bool      `dynamo:"IsActive"`
	CreateAt time.Time `dynamo:"CreateAt"`
	Orders   Orders    `dynamo:"-"`
}

func (s *Sheet) GetPK() *string {
	if s != nil && s.PK != nil {
		return s.PK
	}
	return nil
}

func (s *Sheet) GetSK() *string {
	if s != nil && s.SK != nil {
		return s.SK
	}
	return nil
}

func (s *Sheet) GetName() string {
	if s != nil {
		return s.Name
	}
	return ""
}

func (s *Sheet) GetCoffeeBrand() string {
	if s != nil {
		return s.Brand
	}
	return ""
}

func (s *Sheet) GetMenuURL() string {
	if s != nil {
		return s.MenuURL
	}
	return ""
}

func (s *Sheet) GetHostIDs() string {
	if s != nil {
		return s.HostIDs
	}
	return ""
}

func (s *Sheet) GetIsActive() bool {
	if s != nil {
		return s.IsActive
	}
	return false
}

func (s *Sheet) IsNil() bool {
	if s != nil && s.GetPK() != nil {
		return false
	}
	return true
}

func (s *Sheet) GetModelSheet() *models.Sheet {
	return &models.Sheet{
		SheetID:  s.GetPK(),
		Name:     s.GetName(),
		Brand:    s.GetCoffeeBrand(),
		MenuURL:  s.GetMenuURL(),
		HostIDs:  s.GetHostIDs(),
		IsActive: s.GetIsActive(),
		Orders:   s.Orders.GetModelOrders(),
	}
}

func BuildDomainSheet(v *models.Sheet) *Sheet {
	if v == nil {
		return nil
	}

	return &Sheet{
		PK:       v.SheetID,
		Name:     v.Name,
		Brand:    v.Brand,
		MenuURL:  v.MenuURL,
		HostIDs:  v.HostIDs,
		IsActive: v.IsActive,
		CreateAt: time.Now().UTC(),
	}
}
