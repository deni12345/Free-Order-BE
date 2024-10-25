package logic

// import (
// 	"context"
// 	"fmt"
// 	"github/free-order-be/models"

// 	"log"
// )

// func (l *LogicImpl) CreateSheet(ctx context.Context, req *models.Sheet) (*models.Sheet, error) {
// 	domainSheet := req.BuildDomainSheet()
// 	if domainSheet == nil {
// 		return nil, fmt.Errorf("[Logic] BuilDomainSheet on err nil domain")
// 	}

// 	err := l.Client.SheetDAO.Create(domainSheet)
// 	if err != nil {
// 		log.Printf("[Logic] Create sheet on err: %v", err)
// 		return nil, err
// 	}
// 	return &models.Sheet{}, nil
// }
