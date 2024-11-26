package client

import (
	"net/http"
)

type GetDeliveryIDReq struct {
	URL string `json:"url"`
}

type GetDeliveryIDResp struct {
	Reply struct {
		RestaurantID string `json:"restaurant_id"`
		DeliveryID   string `json:"delivery_id"`
	} `json:"reply"`
}

func (s *ShopeeImpl) GetDeliveryID(req *GetDeliveryIDReq) (*GetDeliveryIDResp, error) {
	url := s.buildURL(Restaurant, req.URL)
	var response *GetDeliveryIDResp
	if err := s.Do(http.MethodGet, url.String(), response); err != nil {
		return nil, err
	}
	return response, nil
}
