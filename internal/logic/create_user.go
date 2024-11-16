package logic

import (
	"context"
	"fmt"

	d "github/free-order-be/internal/domain"
	"github/free-order-be/models"
)

func (l *LogicImpl) CreateUser(ctx context.Context, req *models.User) (*models.User, error) {
	ctxUser := d.BuildDomainUser(req)
	if ctxUser == nil {
		return nil, fmt.Errorf("[Logic] cannot parse model user")
	}

	if user, err := l.Client.UserDAO.FindByEmail(ctx, ctxUser.GetEmail()); err != nil {
		return nil, err
	} else if !user.IsNil() {
		return nil, fmt.Errorf("email %v already exist", ctxUser.GetEmail())
	}

	err := l.Client.UserDAO.Create(ctx, ctxUser)
	if err != nil {
		return nil, err
	}

	return ctxUser.GetModelUser(), nil
}
