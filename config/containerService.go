package config

import (
	"github.com/castillofranciscodaniel/golang-example/pkg/handler"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type ContainerService struct {
	restyClient           *resty.Client
	HealthHandler         handler.HealthHandler
	MessageChannelHandler handler.MessageChannelHandler
}

func NewContainerServiceImp(restyClient *resty.Client, healthHandler handler.HealthHandler, messageChannelHandler handler.MessageChannelHandler) ContainerService {
	configRestyClient(restyClient)

	return ContainerService{
		restyClient:           restyClient,
		HealthHandler:         healthHandler,
		MessageChannelHandler: messageChannelHandler,
	}
}

func configRestyClient(restyClient *resty.Client) {
	restyClient.SetDoNotParseResponse(true)
	restyClient.JSONMarshal = jsoniter.Marshal
	restyClient.JSONUnmarshal = jsoniter.Unmarshal
}
