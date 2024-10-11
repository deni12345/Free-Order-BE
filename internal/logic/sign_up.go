package logic

import (
	"fmt"
	"github/lambda-microservice/internal/domain"
	"github/lambda-microservice/models"
	"log"
)

func (l *LogicImpl) SignUp(req *models.User) (*models.User, error) {
	domainUser := req.BuildDomainUser()
	if domainUser == nil {
		return nil, fmt.Errorf("[Logic] BuildDomainUser on err nil domain")
	}
	err := l.CheckExistedUser(domainUser)
	if err != nil {
		return nil, err
	}
	err = l.Client.UserDAO.Create(domainUser)
	if err != nil {
		log.Printf("[Logic] SignUp on err: %v", err)
		return nil, err
	}

	return models.GetModelUser(domainUser), nil
}

func (l *LogicImpl) CheckExistedUser(userReq *domain.User) error {
	dUser, err := l.Client.UserDAO.Find(userReq)
	if err != nil {
		log.Printf("[Logic] FindUser on err: %v", err)
		return err
	}
	if dUser != nil {
		return fmt.Errorf("[Logic] SignUp on err: User already existed")
	}
	return nil
}
