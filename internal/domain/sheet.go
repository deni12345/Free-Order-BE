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
	ID          *uint     `dynamo:"Id,hash"`
	Name        string    `dynamo:"Name"`
	CoffeeBrand string    `dynamo:"CoffeeBrand"`
	MenuURL     string    `dynamo:"MenuURL"`
	CoopHost    []string  `dynamo:"CoopHost"`
	HostUserID  uint      `dynamo:"HostUserId"`
	IsActive    bool      `dynamo:"IsActive"`
	CreateDatim time.Time `dynamo:"CreateDatim"`
}

func (s *Sheet) GetID() *uint {
	if s != nil {
		return s.ID
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
		return s.CoffeeBrand
	}
	return ""
}

func (s *Sheet) GetMenuURL() string {
	if s != nil {
		return s.MenuURL
	}
	return ""
}

func (s *Sheet) GetCoopHost() []string {
	if s != nil {
		return s.CoopHost
	}
	return nil
}

func (s *Sheet) GetHostUserID() uint {
	if s != nil {
		return s.HostUserID
	}
	return 0
}

func (s *Sheet) GetIsActive() bool {
	if s != nil {
		return s.IsActive
	}
	return false
}

func (s *Sheet) CheckNil() *Sheet {
	if s.ID != nil {
		return s
	}
	return nil
}

func (s *Sheet) GetModelSheet() *models.Sheet {
	return &models.Sheet{
		ID:          s.GetID(),
		Name:        s.GetName(),
		CoffeeBrand: "",
		MenuURL:     "",
		CoopHost:    []string{},
		HostUserID:  0,
		IsActive:    false,
	}
}
