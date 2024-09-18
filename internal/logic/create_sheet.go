package logic

import (
	"context"
	"github/lambda-microservice/models"
)

func (l *LogicImpl) CreateSheet(ctx context.Context, req *models.Sheet) (*models.Sheet, error) {
	return &models.Sheet{}, nil
}
