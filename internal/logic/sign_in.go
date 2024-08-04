package logic

import (
	"context"
	"fmt"
	"github/lambda-microservice/internal/auth"
	"github/lambda-microservice/models"
)

func (l *LogicImpl) SignIn(ctx context.Context, req *models.User) (*models.SignInResp, error) {
	dmu, err := req.BuildDomainUser()
	if err != nil {
		return nil, err
	}
	user, err := l.Client.UserDAO.Find(dmu)
	if err != nil {
		fmt.Printf("[Logic] FindUser on err: %v", err)
		return nil, err
	}
	if !req.ValidIdentity(user) {
		return nil, fmt.Errorf("[Logic] invalid user identity")
	}
	tokenStr, err := auth.CreateToken(user)
	if err != nil {
		return nil, err
	}
	return &models.SignInResp{
		Token:    tokenStr,
		UserName: user.UserName,
		Role:     user.GetRoles(),
	}, nil
}
