package client

import "net/http"

type GetMenuReq struct {
	URL string `json:"url"`
}

type GetMenuResp struct {
	MenuInfos []struct {
		dishes []struct {
			description string `json:"description"`
		} `json:"dishes"`
	} `json:"menu_infos"`
}

func (s *ShopeeImpl) GetMenu(req *GetMenuReq) (*GetMenuResp, error) {
	url, err := s.buildURL("Delivery", req.URL)
	if err != nil {
		return nil, err
	}
	var response *GetMenuResp
	if err := s.Do(http.MethodGet, url, response); err != nil {
		return nil, err
	}
	return response, nil
}
