package client

import "net/http"

type ClientImpl struct {
	EndpointsMap map[string]string
	HTTPClient   *http.Client
}

func NewClientImpl(endpointsMap map[string]string, transport http.RoundTripper) *ClientImpl {
	return &ClientImpl{
		EndpointsMap: endpointsMap,
		HTTPClient: &http.Client{
			Transport: transport,
		},
	}
}

type ShopeeTransport struct {
	transport http.Transport
	header    map[string]string
}

func (s *ShopeeTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	return nil, nil
}
