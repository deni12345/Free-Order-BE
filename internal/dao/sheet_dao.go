package dao

import (
	"context"
	"fmt"
	d "github/free-order-be/internal/domain"

	"github.com/guregu/dynamo/v2"
)

type ISheetDAO interface {
	CreateInfo(context.Context, *d.Sheet) error
	FindsByName(context.Context, string) (d.Sheets, error)
}

type SheetImpl struct {
	dao   *DAO
	table dynamo.Table
}

func NewSheetDAO(db *dynamo.DB) *SheetImpl {
	return &SheetImpl{
		dao:   NewDAO(db),
		table: db.Table(SHEET_TABLE),
	}
}

func (s *SheetImpl) CreateInfo(ctx context.Context, sheet *d.Sheet) error {
	newID, err := s.dao.NextID(ctx, SHEET_TABLE)
	if err != nil {
		return err
	}
	if newID == nil {
		return fmt.Errorf("failed to get next id")
	}
	sheet.PK = s.createSheetPK(newID)
	sheet.SK = SHEET_SK
	return s.table.Put(sheet).Run(ctx)
}

func (s *SheetImpl) FindsByName(ctx context.Context, name string) (d.Sheets, error) {
	var sheets d.Sheets
	err := s.table.Scan().Filter("'Name'=?", name).All(ctx, &sheets)
	if err != nil {
		return nil, err
	}
	return sheets, nil
}

func (s *SheetImpl) createSheetPK(id *uint) string {
	return fmt.Sprintf("SHEET#%v", *id)
}
