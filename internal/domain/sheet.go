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
	PK          string    `dynamo:"PK,hash"`
	SK          string    `dynamo:"SK,range"`
	Name        string    `dynamo:"Name"`
	Brand       string    `dynamo:"Brand"`
	MenuURL     string    `dynamo:"MenuURL"`
	HostIDs     string    `dynamo:"HostUserId"`
	IsActive    bool      `dynamo:"IsActive"`
	CreateDatim time.Time `dynamo:"CreateDatim"`
}

func (s *Sheet) GetID() string {
	if s != nil {
		return s.PK
	}
	return ""
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

func (s *Sheet) GetSK() string {
	if s != nil {
		return s.SK
	}
	return ""
}

func (s *Sheet) CheckNil() *Sheet {
	if s.PK != "" {
		return s
	}
	return nil
}

func (s *Sheet) GetModelSheet() *models.Sheet {
	return &models.Sheet{
		PK:       s.GetID(),
		Name:     s.GetName(),
		Brand:    s.GetCoffeeBrand(),
		MenuURL:  s.GetMenuURL(),
		HostIDs:  s.GetHostIDs(),
		IsActive: s.GetIsActive(),
	}
}

func BuildDomainSheet(v *models.Sheet) *Sheet {
	if v == nil {
		return nil
	}
	SK := "INFO#METADATA"
	return &Sheet{
		PK:          v.PK,
		SK:          SK,
		Name:        v.Name,
		Brand:       v.Brand,
		MenuURL:     v.MenuURL,
		HostIDs:     v.HostIDs,
		IsActive:    v.IsActive,
		CreateDatim: time.Now().UTC(),
	}
}
