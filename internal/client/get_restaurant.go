package client

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type GetDeliveryIDReq struct {
	URL string `json:"url"`
}

type GetDeliveryIDRes struct {
	DeliveryID uint `json:"id"`
}

func (c *ClientImpl) GetDeliveryID(req *GetDeliveryIDReq) (*GetDeliveryIDRes, error) {
	// endpoint, ok := c.EndpointsMap["DeliveryID"]
	// if !ok {
	// 	return nil, fmt.Errorf("endpoint not found")
	// }

	request, err := http.NewRequest("GET", "https://gappapi.deliverynow.vn/api/delivery/get_from_url?url=ho-chi-minh/fruit-crush-trai-cay-tuoi-an-vat-thich-quang-duc", nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-Foody-Access-Token", "")
	request.Header.Add("X-Foody-Api-Version", "1")
	request.Header.Add("X-Foody-App-Type", "1004")
	request.Header.Add("X-Foody-Client-Id", "")
	request.Header.Add("X-Foody-Client-Language", "vi")
	request.Header.Add("X-Foody-Client-Type", "1")
	request.Header.Add("X-Foody-Client-Version", "3.0.0")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	logrus.Infof("response: %v", string(body))
	return nil, nil
}
