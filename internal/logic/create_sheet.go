package logic

import (
	"context"
	"fmt"
	d "github/free-order-be/internal/domain"
	"github/free-order-be/models"

	"log"
)

func (l *LogicImpl) CreateSheet(ctx context.Context, req *models.Sheet) (*models.Sheet, error) {
	ctxSheet := d.BuildDomainSheet(req)
	if ctxSheet == nil {
		return nil, fmt.Errorf("[Logic] cannot parse model user")
	}
	sheets, err := l.Client.SheetDAO.FindsByName(ctx, ctxSheet.GetName())
	if err != nil {
		return nil, err
	}
	if len(sheets) > 0 {
		return nil, fmt.Errorf("sheet %v already exist", ctxSheet.GetName())
	}

	err = l.Client.SheetDAO.CreateInfo(ctx, ctxSheet)
	if err != nil {
		log.Printf("[Logic] Create sheet on err: %v", err)
		return nil, err
	}
	return ctxSheet.GetModelSheet(), nil
}
