package dao

import (
	"context"
	"fmt"
	d "github/free-order-be/internal/domain"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/guregu/dynamo/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type ISheetDAO interface {
	CreateInfo(context.Context, *d.Sheet) error
	FindByID(context.Context, string) (*d.Sheet, error)
	FindsByName(context.Context, string) (d.Sheets, error)
}

type SheetImpl struct {
	dao   *DAO
	table dynamo.Table
}

func NewSheetDAO(db *dynamo.DB) *SheetImpl {
	return &SheetImpl{
		dao:   NewDAORef(db),
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
	sheet.PK = createSheetPK(newID)
	sheet.SK = aws.String(SHEET_SK)
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

func (s *SheetImpl) FindByID(ctx context.Context, id string) (*d.Sheet, error) {
	var result = []map[string]types.AttributeValue{}
	err := s.table.Scan().Filter("PK=?", id).All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return toSheet(result), nil
}

// This convert dynamo sheet fields to sheet domain which contains orders in go
func toSheet(items []map[string]types.AttributeValue) *d.Sheet {
	var (
		mu    sync.Mutex
		sheet = &d.Sheet{}
		eg    = errgroup.Group{}
	)
	for _, item := range items {
		item := item
		eg.Go(func() error {
			SK, ok := item["SK"].(*types.AttributeValueMemberS)
			if !ok {
				return fmt.Errorf("failed to get SK")
			}
			if strings.Contains(SK.Value, "ORDER#") {
				var order *d.Order
				if err := dynamo.UnmarshalItem(item, &order); err != nil {
					return err
				}
				mu.Lock()
				defer mu.Unlock()
				sheet.Orders = append(sheet.Orders, order)
				return nil

			}
			return dynamo.UnmarshalItem(item, sheet)
		})
	}
	if err := eg.Wait(); err != nil {
		logrus.Infof("error unmarshal sheet: %v", err)
		return nil
	}
	return sheet
}

func createSheetPK(id *uint) *string {
	return aws.String(fmt.Sprintf("SHEET#%v", *id))
}
