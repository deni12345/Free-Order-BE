package logic

import (
	"context"
	"fmt"
	"github/free-order-be/internal/auth"
	"github/free-order-be/models"
)

func (l *LogicImpl) SignIn(ctx context.Context, req *models.User) (*models.SignInResp, error) {
	domainUser := req.BuildDomainUser()
	if domainUser == nil {
		return nil, fmt.Errorf("[Logic] BuildDomainUser on err nil domain")
	}
	user, err := l.Client.UserDAO.Find(domainUser)
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
