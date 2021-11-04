package config

import (
	"github.com/castillofranciscodaniel/golang-example/pkg/handler"
	"github.com/go-resty/resty/v2"
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
}
