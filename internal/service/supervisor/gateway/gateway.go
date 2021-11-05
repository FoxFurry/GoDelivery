package gateway

import (
	"bytes"
	"encoding/json"
	"github.com/foxfurry/go_delivery/internal/domain/dto"
	"github.com/spf13/viper"
	"net/http"
)

type IGateway interface {
	GetMenu() (*http.Response, error)
	SendOrder(order dto.Order) (*http.Response, error)
	Distribute(distro dto.Distribution) (*http.Response, error)
}

type deliveryGateway struct {}

func NewDeliveryGateway() IGateway {
	return &deliveryGateway{}
}

func (g *deliveryGateway) GetMenu() (*http.Response, error) {
	return http.Get(viper.GetString("aggregator_host") + "/menu")
}

func (g *deliveryGateway) SendOrder(order dto.Order) (*http.Response, error) {
	jsonBody, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	contentType := "application/json"
	return http.Post(viper.GetString("aggregator_host") + "/order", contentType, bytes.NewReader(jsonBody))
}

func (g *deliveryGateway) Distribute(distro dto.Distribution) (*http.Response, error) {
	jsonBody, err := json.Marshal(distro)
	if err != nil {
		return nil, err
	}

	contentType := "application/json"
	return http.Post(viper.GetString("client_host") + "/distribution", contentType, bytes.NewReader(jsonBody))

}

