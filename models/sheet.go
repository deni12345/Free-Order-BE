package models

// import (
// 	"github/free-order-be/internal/domain"
// 	"time"
// )

// type Sheet struct {
// 	ID        *uint      `json:"id"`
// 	UserID    *uint      `json:"user_id"`
// 	Name      string     `json:"name"`
// 	EndAt     *time.Time `json:"end_at"`
// 	CreatedAt *time.Time `json:"created_at,omitempty"`
// 	IsActive  bool       `json:"is_active"`
// 	Orders    Orders     `json:"orders"`
// }

// type Orders []*Order

// type Order struct {
// 	ID       *uint  `json:"id"`
// 	OrderBy  string `json:"order_by"`
// 	FoodName string `json:"food_name"`
// 	Amount   int32  `json:"amount"`
// 	Note     string `json:"note"`
// 	CreateAt string `json:"create_at"`
// }

// func (sh *Sheet) GetID() *uint {
// 	if sh != nil {
// 		return sh.ID
// 	}
// 	return nil
// }

// func (sh *Sheet) GetName() string {
// 	if sh != nil {
// 		return sh.Name
// 	}
// 	return ""
// }

// func (sh *Sheet) GetCreatedBy() string {
// 	if sh != nil {
// 		return sh.Name
// 	}
// 	return ""
// }

// func (sh *Sheet) GetCreatedAt() *time.Time {
// 	if sh != nil {
// 		return sh.CreatedAt
// 	}
// 	return nil
// }

// func (sh *Sheet) GetUserID() *uint {
// 	if sh != nil {
// 		return sh.UserID
// 	}
// 	return nil
// }

// func (sh *Sheet) GetOrders() Orders {
// 	if sh != nil {
// 		return sh.Orders
// 	}
// 	return nil
// }

// func (sh *Sheet) GetIsActive() bool {
// 	if sh != nil {
// 		return sh.IsActive
// 	}
// 	return true
// }

// func (sh *Sheet) BuildDomainSheet() *domain.Sheet {
// 	if sh == nil {
// 		return nil
// 	}
// 	return &domain.Sheet{
// 		ID:       sh.GetID(),
// 		Name:     sh.GetName(),
// 		UserId:   sh.GetUserID(),
// 		IsActive: sh.GetIsActive(),
// 		//Order:  sh.GetOrders(),
// 	}
// }
