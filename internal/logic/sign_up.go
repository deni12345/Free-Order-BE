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
	err := l.CheckExistedUser(ctx, ctxUser)
	if err != nil {
		return nil, err
	}
	err = l.Client.UserDAO.Create(ctx, ctxUser)
	if err != nil {
		log.Printf("[Logic] SignUp on err: %v", err)
		return nil, err
	}

	return ctxUser.GetModelUser(), nil
}

func (l *LogicImpl) CheckExistedUser(ctx context.Context, req *d.User) error {
	dbUser, err := l.Client.UserDAO.FindByEmail(ctx, req.GetEmail())
	if err != nil {
		log.Printf("[Logic] find user on err: %v", err)
		return err
	}
	if len(dbUser) >= 1 {
		return fmt.Errorf("[Logic] user already existed")
	}
	return nil
}
