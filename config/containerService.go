package config

import (
	"github.com/castillofranciscodaniel/golang-example/pkg/handler"
)

type ContainerServiceImp struct {
	HealthHandler  handler.HealthHandler
	ProductHandler handler.ProductHandler
}

func NewContainerServiceImp(healthHandler handler.HealthHandler, productHandler handler.ProductHandler) ContainerServiceImp {
	return ContainerServiceImp{
		HealthHandler:  healthHandler,
		ProductHandler: productHandler,
	}
}
