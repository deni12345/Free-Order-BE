package logic

import (
	"context"
	"fmt"

	d "github/free-order-be/internal/domain"
	"github/free-order-be/models"
	"log"
)

func (l *LogicImpl) SignUp(ctx context.Context, req *models.User) (*models.User, error) {
	ctxUser := d.BuildDomainUser(req)
	if ctxUser == nil {
		return nil, fmt.Errorf("[Logic] cannot parse model user")
	}
	users, err := l.Client.UserDAO.FindByEmail(ctx, ctxUser.GetEmail())
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		return nil, fmt.Errorf("email %v already exist", ctxUser.GetEmail())
	}

	err = l.Client.UserDAO.Create(ctx, ctxUser)
	if err != nil {
		log.Printf("[Logic] SignUp on err: %v", err)
		return nil, err
	}

	return ctxUser.GetModelUser(), nil
}
