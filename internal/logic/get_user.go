package logic

import (
	"context"
	"fmt"
	"github/free-order-be/models"
)

func (l *LogicImpl) GetUser(ctx context.Context, req *models.GetUserReq) (*models.User, error) {
	if req == nil {
		return nil, fmt.Errorf("[Logic] BuildDomainUser on err nil domain")
	}
	result, err := l.Client.UserDAO.FindByID(ctx, req.GetUserID())
	if err != nil {
		return nil, err
	}
	return result.GetModelUser(), nil
}
