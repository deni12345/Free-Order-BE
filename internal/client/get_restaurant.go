package client

import (
	"net/http"
	"net/url"
)

type GetDeliveryIDReq struct {
	RestaurantEndpoint string
}

type GetRestaurantIDResp struct {
	RestaurantID uint `json:"restaurant_id"`
	DeliveryID   uint `json:"delivery_id"`
}

func (res *GetRestaurantIDResp) GetRestaurantID() uint {
	if res != nil {
		return res.RestaurantID
	}
	return 0
}

func (s *ShopeeImpl) GetRestaurantID(req *GetDeliveryIDReq) (*GetRestaurantIDResp, error) {
	url := s.buildURL(Restaurant, req.toQuery())
	var response *GetRestaurantIDResp
	if err := s.Do(http.MethodGet, url.String(), response); err != nil {
		return nil, err
	}
	return response, nil
}

func (req *GetDeliveryIDReq) toQuery() url.Values {
	q := url.Values{}
	q.Set("url", req.RestaurantEndpoint)
	return q
}
