package models

import "github/lambda-microservice/internal/domain"

type User struct {
	UserID   *int64 `json: "user_id"`
	UserName string `json: "user_name"`
	Password string `json: "pasword"`
}

func (mUser *User) BuildDomainUser() (*domain.User, error) {
	hashPasword, err := hashPassword(mUser.Password)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		UserName: mUser.UserName,
		Password: hashPasword,
	}, nil
}
