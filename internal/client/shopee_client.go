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
		Dishes:     "/dish/get_delivery_dishes?id_type=%s&request_id=%s",
		Restaurant: "/delivery/get_from_url?url=%s",
	}
)

type Shopee interface {
	GetRestaurant(*GetDeliveryIDReq) (*GetDeliveryIDResp, error)
}

type ShopeeImpl struct {
	EndpointsMap map[Endpoint]string
	HTTPClient   *http.Client
}

func NewClientImpl() *ShopeeImpl {
	return &ShopeeImpl{
		EndpointsMap: endpointsMap,
		HTTPClient: &http.Client{
			Transport: NewShopeeTransport(),
		},
	}
}

func NewShopeeTransport() *ShopeeTransport {
	header := map[string]string{
		"X-Foody-Access-Token":    "",
		"X-Foody-Api-Version":     "1",
		"X-Foody-App-Type":        "1004",
		"X-Foody-Client-Id":       "",
		"X-Foody-Client-Language": "vi",
		"X-Foody-Client-Type":     "1",
		"X-Foody-Client-Version":  "3.0.0",
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

func (s *ShopeeImpl) buildURL(endpoint Endpoint, urlPath string) (string, error) {
	return url.JoinPath(baseShopeeURL, s.EndpointsMap[endpoint], urlPath)
}

func (s *ShopeeImpl) Do(method string, url string, resp interface{}) error {
	httpReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	htppResp, err := s.HTTPClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer htppResp.Body.Close()
	if err = json.NewDecoder(htppResp.Body).Decode(resp); err != nil {
		return err
	}
	return nil
}
