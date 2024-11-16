package logic

import (
	"context"
	"fmt"
	"github/free-order-be/internal/auth"
	"github/free-order-be/internal/domain"
	"github/free-order-be/models"
)

func (l *LogicImpl) SignIn(ctx context.Context, req *models.User) (*models.SignInResp, error) {
	ctxUser := domain.BuildDomainUser(req)
	if ctxUser == nil || ctxUser.GetEmail() == "" {
		return nil, fmt.Errorf("[Logic] invalid user")
	}
	user, err := l.Client.UserDAO.FindByEmail(ctx, ctxUser.GetEmail())
	if err != nil {
		fmt.Printf("[Logic] find user by email on err: %v", err)
		return nil, err
	}
	if user.IsNil() {
		return nil, fmt.Errorf("[Logic] user not found")
	}
	if !user.IsValid(ctxUser) {
		return nil, fmt.Errorf("[Logic] invalid user")
	}

	tokenStr, err := auth.CreateToken(user.GetModelUser())
	if err != nil {
		return nil, err
	}
	return &models.SignInResp{
		Token:    tokenStr,
		UserName: user.GetName(),
		Email:    user.GetEmail(),
	}, nil
}
