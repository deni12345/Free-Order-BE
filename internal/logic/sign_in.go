package logic

import (
	"context"
	"github/lambda-microservice/models"
	"log"
)

func (l *LogicImpl) SignIn(ctx context.Context, req *models.User) (*models.SignInResp, error) {
	dmu, err := req.BuildDomainUser()
	if err != nil {
		return nil, err
	}
	err = l.CheckExistedUser(dmu)
	if err != nil {
		return nil, err
	}
	err = l.Client.UserDAO.Create(dmu)
	if err != nil {
		log.Printf("Logic CreateUser on err: %v", err)
		return nil, err
	}

	return nil, nil
}
