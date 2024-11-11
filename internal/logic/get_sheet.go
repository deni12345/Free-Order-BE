package logic

import (
	"context"
	"fmt"
	"github/free-order-be/models"
)

func (l *LogicImpl) GetSheet(ctx context.Context, req *models.GetSheetReq) (*models.Sheet, error) {
	if req == nil {
		return nil, fmt.Errorf("[Logic] Invalid get sheet request")
	}
	sheet, err := l.Client.SheetDAO.FindByID(ctx, req.GetSheetID())
	if err != nil {
		return nil, err
	}
	if sheet.IsNil() {
		return nil, fmt.Errorf("[Logic] there is no sheet founded")
	}
	return sheet.GetModelSheet(), nil
}
