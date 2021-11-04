package config

import (
	"github.com/castillofranciscodaniel/golang-example/pkg/handler"
)

type ContainerServiceImp struct {
	HealthHandler         handler.HealthHandler
	MessageChannelHandler handler.MessageChannelHandler
}

func NewContainerServiceImp(healthHandler handler.HealthHandler, messageChannelHandler handler.MessageChannelHandler) ContainerServiceImp {
	return ContainerServiceImp{
		HealthHandler:         healthHandler,
		MessageChannelHandler: messageChannelHandler,
	}
}
