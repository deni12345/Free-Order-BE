package client

import (
	"fmt"
	"github/free-order-be/internal/client/models"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/sync/errgroup"
)

type GetDishesReq struct {
	RestaurantID    uint `json:"restaurant_id"`
	EnableUniMartAb int  `json:"enable_uni_mart_ab"`
}

type GetDishesResp struct {
	MenuDishes []*MenuDish
}

type MenuDish struct {
	ID            uint
	Name          string
	PictureUrlFmt string
	TotalLike     int
	Price         float64
}

func (s *ShopeeImpl) GetDishes(req *GetDishesReq) (*GetDishesResp, error) {
	url := s.buildURL(Dishes, req.toQuery())
	var res struct {
		Data *models.ShopeeDishesResp `json:"data"`
	}
	if err := s.Do(http.MethodGet, url.String(), &res); err != nil {
		return nil, err
	}
	return &GetDishesResp{toMenuDishes(res.Data)}, nil
}

func toMenuDishes(resp *models.ShopeeDishesResp) (menu []*MenuDish) {
	var (
		erg errgroup.Group
		mu  sync.Mutex
	)
	for _, catalog := range resp.Catalogs {
		for _, dish := range catalog.Dishes {
			erg.Go(func() error {
				mu.Lock()
				defer mu.Unlock()
				menu = append(menu, getMenuDish(dish))
				return nil
			})
		}
	}
	if err := erg.Wait(); err != nil {
		return nil
	}
	return menu
}

func getMenuDish(dish *models.Dish) *MenuDish {
	return &MenuDish{
		dish.ID,
		dish.Name,
		dish.PictureUrlFmt,
		dish.TotalLike,
		dish.Price,
	}
}

func (req *GetDishesReq) toQuery() url.Values {
	q := url.Values{}
	q.Set("restaurant_id", fmt.Sprintf("%d", req.RestaurantID))
	q.Set("enable_uni_mart_ab", fmt.Sprintf("%d", req.EnableUniMartAb))
	return q
}
