package client

import (
	m "github/free-order-be/internal/client/models"
	"net/http"
)

type GetDishesReq struct {
	URL string `json:"url"`
}

type GetDishesResp struct {
	Data struct {
		Catalogs m.Catalogs `json:"catalogs"`
	} `json:"data"`
}

func (s *ShopeeImpl) GetDishes(req *GetDishesReq) (*GetDishesResp, error) {
	url := s.buildURL(Dishes, req.URL)
	var response *GetDishesResp
	if err := s.Do(http.MethodGet, url.String(), response); err != nil {
		return nil, err
	}
	return response, nil
}
