package domain

import (
	"github/free-order-be/models"
	"time"
)

type Users []*User
type User struct {
	ID           *uint     `dynamo:"Id,hash"`
	UserName     string    `dynamo:"Name,range"`
	HashPassword string    `dynamo:"Password"`
	Email        string    `dynamo:"Email"`
	Phone        string    `dynamo:"Phone"`
	CreateDatim  time.Time `dynamo:"CreateDatim"`
}

func (u *User) GetID() *uint {
	if u != nil {
		return u.ID
	}
	return nil
}

func (u *User) GetName() string {
	if u != nil {
		return u.UserName
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
		return u.UserName
	}
	return ""
}

func (u *User) GetCreateDatim() time.Time {
	if u != nil {
		return u.CreateDatim
	}
	return time.Time{}
}

func (u *User) CheckNil() *User {
	if u.ID != nil {
		return u
	}
	return nil
}

func (u *User) GetModelUser() *models.User {
	return &models.User{
		ID:          u.GetID(),
		Email:       u.GetEmail(),
		Phone:       u.GetPhone(),
		UserName:    u.GetName(),
		Password:    u.GetHashPassword(),
		CreateDatim: u.GetCreateDatim(),
	}
}

func BuildDomainUser(v *models.User) *User {
	if v == nil {
		return nil
	}
	hashPasword, err := HashPassword(v.Password)
	if err != nil {
		return nil
	}
	return &User{
		ID:           v.ID,
		UserName:     v.UserName,
		HashPassword: hashPasword,
		Email:        v.Email,
		Phone:        v.Phone,
		CreateDatim:  time.Now().UTC(),
	}
}

func (u *User) ValidIdentity(v *models.User) bool {
	if u != nil {
		return CheckPasswordHash(v.Password, u.HashPassword)
	}
	return false
}
