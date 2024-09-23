package logic

import (
	"context"
	"fmt"
	"github/lambda-microservice/models"
	"log"
)

func (l *LogicImpl) CreateSheet(ctx context.Context, req *models.Sheet) (*models.Sheet, error) {
	domainSheet := req.BuildDomainSheet()
	if domainSheet == nil {
		return nil, fmt.Errorf("[Logic] BuilDomainSheet on err nil domain")
	}

	err = l.Client.UserDAO.Create(domainUser)
	if err != nil {
		log.Printf("[Logic] SignUp on err: %v", err)
		return nil, err
	}
	return &models.Sheet{}, nil
}
