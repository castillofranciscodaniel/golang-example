//+build InitializeServer

package main

import (
	configDb "bitbucket.org/chattigodev/chattigo-golang-ig-config-library/pkg/config"
	coreConfig "bitbucket.org/chattigodev/chattigo-golang-ig-core-library/pkg/dbCore/config"
	"bitbucket.org/chattigodev/chattigo-golang-kafka-library/pkg/producer"
	"github.com/castillofranciscodaniel/golang-example/config"
	"github.com/castillofranciscodaniel/golang-example/pkg/client"
	"github.com/castillofranciscodaniel/golang-example/pkg/handler"
	"github.com/castillofranciscodaniel/golang-example/pkg/service"
	"github.com/google/wire"
)

func InitializeServer() config.ContainerServiceImp {
	wire.Build(
		NewDeserializerServiceImpl,
		NewSerializerServiceImpl,
		producer.NewProducer,
		configDb.NewIgDbConfigImpl,
		coreConfig.NewIgDbCoreImpl,

		client.NewProductClientImpl,
		service.NewProductServiceIml,
		handler.NewProductHandler,

		config.NewContainerServiceImp,
	)
	return config.ContainerServiceImp{}
}
