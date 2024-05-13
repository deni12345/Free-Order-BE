package logic

import (
	"fmt"
	"github/lambda-microservice/internal/domain"
	"github/lambda-microservice/models"
	"log"
)

func (l *LogicImpl) CreateUser(req *models.User) (*models.User, error) {
	userReq, err := req.BuildDomainUser()
	if err != nil {
		log.Printf("Logic build domain user on err: %v", err)
		return nil, err
	}
	err = l.CheckExistedUser(userReq)
	if err != nil {
		log.Printf("Logic check existed user on err: %v", err)
		return nil, err
	}
	err = l.Client.UserDAO.Create(userReq)
	if err != nil {
		return nil, err
	}
	return &models.User{
		UserName: userReq.UserName,
		Password: userReq.Password,
	}, nil
}

func (l *LogicImpl) CheckExistedUser(userReq *domain.User) error {
	dUser, err := l.Client.UserDAO.Find(userReq)
	if err != nil {
		return err
	}
	if dUser != nil {
		return fmt.Errorf("Logic create user on err: User already existed")
	}
	return nil
}
