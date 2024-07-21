package models

import (
	"github/lambda-microservice/internal/domain"
	"log"
)

type User struct {
	UserID   *int64 `json:"user_id,omitempty"`
	UserName string `json:"user_name"`
	Password string `json:"password,omitempty"`
}

type FindUserResp struct {
	Token    string `json:"token"`
	UserID   int64  `json:"user_id,omitempty"`
	UserName string `json:"user_name"`
}

func (mu *User) BuildDomainUser() (*domain.User, error) {
	hashPasword, err := HashPassword(mu.Password)
	if err != nil {
		log.Printf("Logic build domain user on err: %v", err)
		return nil, err
	}
	return &domain.User{
		UserName: mu.UserName,
		Password: hashPasword,
	}, nil
}

func GetModelUser(dmu *domain.User) *User {
	return &User{
		UserName: dmu.UserName,
		Password: dmu.Password,
	}
}
