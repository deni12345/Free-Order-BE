package logic

import (
	"context"
	"github/free-order-be/internal/client"
)

func (l *LogicImpl) GetShopeeMenu(ctx context.Context) (*client.GetDeliveryIDRes, error) {
	clientIns := client.NewClient(map[string]string{})
	res, err := clientIns.GetDeliveryID(&client.GetDeliveryIDReq{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
