package domain

import (
	"github/free-order-be/models"
	"time"

	"github.com/sirupsen/logrus"
)

type Users []*User

type User struct {
	ID           *string   `dynamo:"ID,hash"`
	Name         string    `dynamo:"Name,range"`
	HashPassword string    `dynamo:"Password"`
	Email        string    `dynamo:"Email"`
	Phone        string    `dynamo:"Phone"`
	CreateDatim  time.Time `dynamo:"CreateDatim"`
	GoogleID     string    `dynamo:"GoogleID"`
	IsActive     bool      `dynamo:"IsActive"`
}

func (u *User) GetID() *string {
	if u != nil {
		return u.ID
	}
	return nil
}

func (u *User) GetName() string {
	if u != nil {
		return u.Name
	}
	return ""
}

func (u *User) GetHashPassword() string {
	if u != nil {
		return u.HashPassword
	}
	return ""
}

func (u *User) GetEmail() string {
	if u != nil {
		return u.Email
	}
	return ""
}

func (u *User) GetPhone() string {
	if u != nil {
		return u.Name
	}
	return ""
}

func (u *User) GetCreateDatim() time.Time {
	if u != nil {
		return u.CreateDatim
	}
	return time.Time{}
}

func (u *User) GetIsActive() bool {
	if u != nil {
		return u.IsActive
	}
	return false
}

func (u *User) IsNil() bool {
	if u != nil && u.ID != nil {
		return false
	}
	return true
}

func (u *User) GetModelUser() *models.User {
	return &models.User{
		ID:          u.GetID(),
		Email:       u.GetEmail(),
		Phone:       u.GetPhone(),
		UserName:    u.GetName(),
		Password:    u.GetHashPassword(),
		IsActive:    u.GetIsActive(),
		CreateDatim: u.GetCreateDatim(),
	}
}

func BuildDomainUser(v *models.User) *User {
	if v == nil {
		return nil
	}
	logrus.Infof("Test %v", v.Password)
	hashPasword, err := HashPassword(v.Password)
	if err != nil {
		return nil
	}
	return &User{
		ID:           v.ID,
		Name:         v.UserName,
		HashPassword: hashPasword,
		Email:        v.Email,
		Phone:        v.Phone,
		IsActive:     v.IsActive,
		GoogleID:     v.GoogleID,
		CreateDatim:  time.Now().UTC(),
	}
}

func (u *User) IsValid(v *User) bool {
	if u != nil {
		return CheckPasswordHash(v.GetHashPassword(), u.GetHashPassword())
	}
	return false
}
