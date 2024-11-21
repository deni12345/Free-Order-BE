package logic

import (
	"context"
	"github/free-order-be/internal/client"
)

func (l *LogicImpl) GetShopeeMenu(ctx context.Context, url string) (*client.GetDeliveryIDResp, error) {

	res, err := clientIns.GetDeliveryID(&client.GetDeliveryIDReq{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
