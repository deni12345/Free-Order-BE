package domain

import (
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

func (u *User) CheckNil() *User {
	if u.ID != nil {
		return u
	}
	return nil
}
