package dao

import (
	"context"
	"fmt"
	d "github/free-order-be/internal/domain"

	"github.com/guregu/dynamo/v2"
)

type ISheetDAO interface {
	CreateInfo(context.Context, *d.Sheet) error
	Find(context.Context, string) (d.Sheets, error)
}

type SheetImpl struct {
	dao   *DAO
	table dynamo.Table
}

func (s *SheetImpl) CreateInfo(ctx context.Context, sheet *d.Sheet) error {
	newID, err := s.dao.NextID(ctx, SHEET_TABLE)
	if err != nil {
		return err
	}

	sheet.PK = fmt.Sprintf("SHEET#%v", *newID)
	sheet.SK = "INFO#METADATA"
	return s.table.Put(sheet).Run(ctx)
}

func (s *SheetImpl) Find(ctx context.Context, name string) (d.Sheets, error) {
	var results = d.Sheets{}
	err := s.table.Scan().Filter("'Name'=?", name).All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
