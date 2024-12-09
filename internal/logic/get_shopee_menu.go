package logic

import (
	"context"
	"fmt"
	"github/free-order-be/internal/client"
)

func (l *LogicImpl) GetShopeeMenu(ctx context.Context, endpoint string) (*client.GetDishesResp, error) {
	res, err := l.Shopee.GetRestaurantID(&client.GetDeliveryIDReq{
		RestaurantEndpoint: endpoint,
	})
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("restaurant endpoint [%v] not found", endpoint)
	}
	menu, err := l.Shopee.GetDishes(&client.GetDishesReq{
		RestaurantID: res.GetRestaurantID(),
	})
	if err != nil {
		return nil, err
	}
	if menu == nil {
		return nil, fmt.Errorf("menu not found: %v", endpoint)
	}
	return menu, nil
}
