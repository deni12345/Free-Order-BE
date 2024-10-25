package logic

import (
	"context"
	"github/free-order-be/models"
)

func (l *LogicImpl) GetUser(ctx context.Context, req *models.GetUserReq) (*models.User, error) {
	// if req == nil {
	// 	return nil, fmt.Errorf("[Logic] BuildDomainUser on err nil domain")
	// }
	// result, err := l.Client.UserDAO.FindByEmail(ctx, req.GetUserEmail())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}
