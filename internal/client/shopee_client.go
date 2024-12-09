package client

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Endpoint int

const (
	Dishes Endpoint = iota
	Restaurant
)

var (
	baseShopeeURL = "https://gappapi.deliverynow.vn/api"
	endpointsMap  = map[Endpoint]string{
		Dishes:     "/v6/buyer/store/dishes",
		Restaurant: "/delivery/get_from_url",
	}
)

type Shopee interface {
	GetDishes(*GetDishesReq) (*GetDishesResp, error)
	GetRestaurantID(*GetDeliveryIDReq) (*GetRestaurantIDResp, error)
}

type ShopeeImpl struct {
	EndpointsMap map[Endpoint]string
	HttpClient   *http.Client
}

func NewClientImpl() *ShopeeImpl {
	return &ShopeeImpl{
		EndpointsMap: endpointsMap,
		HttpClient: &http.Client{
			Transport: NewShopeeTransport(),
		},
	}
}

func NewShopeeTransport() *ShopeeTransport {
	header := map[string]string{
		"x-foody-client-id":             "",
		"x-foody-client-language":       "vi",
		"x-foody-api-version":           "1",
		"x-foody-app-type":              "3000",
		"x-foody-client-type":           "3",
		"X-Shopee-Client-Timezone":      "Asia/Ho_Chi_Minh",
		"x-foody-client-version":        "3.38.28",
		"x-foody-client-bundle-version": "603407",
	}

	return &ShopeeTransport{
		transport: http.DefaultTransport,
		header:    header,
	}
}

type ShopeeTransport struct {
	transport http.RoundTripper
	header    map[string]string
}

func (s *ShopeeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range s.header {
		req.Header.Set(k, v)
	}
	return s.transport.RoundTrip(req)
}

func (s *ShopeeImpl) buildURL(endpoint Endpoint, query url.Values) url.URL {
	url := url.URL{
		Scheme: "https",
		Host:   baseShopeeURL,
		Path:   s.EndpointsMap[endpoint],
	}

	if query != nil {
		url.ForceQuery = true
		url.RawQuery = query.Encode()
	}
	return url
}

func (s *ShopeeImpl) Do(method string, url string, resp interface{}) error {
	httpReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	htppResp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer htppResp.Body.Close()
	if err = json.NewDecoder(htppResp.Body).Decode(resp); err != nil {
		return err
	}
	return nil
}
