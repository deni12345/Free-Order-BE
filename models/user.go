package models

import (
	"github/lambda-microservice/internal/domain"
	"log"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password,omitempty"`
}

type SignInResp struct {
	Token    string   `json:"token"`
	Role     []string `json:"role"`
	UserName string   `json:"user_name"`
}

func (mu *User) BuildDomainUser() (*domain.User, error) {
	hashPasword, err := HashPassword(mu.Password)
	if err != nil {
		log.Printf("Logic build domain user on err: %v", err)
		return nil, err
	}
	return &domain.User{
		UserName:     mu.UserName,
		HashPassword: hashPasword,
	}, nil
}

func GetModelUser(dmu *domain.User) *User {
	return &User{
		UserName: dmu.UserName,
		Password: dmu.HashPassword,
	}
}

func (u *User) ValidIdentity(dmu *domain.User) bool {
	if u != nil {
		return CheckPasswordHash(u.Password, dmu.HashPassword)
	}
	return false
}
