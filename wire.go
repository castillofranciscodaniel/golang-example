//+build InitializeServer

package main

import (
	service "bitbucket.org/chattigodev/chattigo-golang-library/pkg/deserializer"
	"bitbucket.org/chattigodev/chattigo-golang-library/pkg/http"
	service2 "bitbucket.org/chattigodev/chattigo-golang-library/pkg/serializer"
	"github.com/castillofranciscodaniel/golang-example/config"
	"github.com/castillofranciscodaniel/golang-example/pkg/client"
	"github.com/castillofranciscodaniel/golang-example/pkg/handler"
	service3 "github.com/castillofranciscodaniel/golang-example/pkg/service"
	"github.com/google/wire"
)

func InitializeServer() config.ContainerServiceImp {
	wire.Build(
		service.NewDeserializerServiceImpl,
		service2.NewSerializerServiceImpl,
		http.NewWebClientImpl,

		client.NewProductClientImpl,
		service3.NewProductServiceIml,
		handler.NewHealthHandler,
		handler.NewProductHandler,

		config.NewContainerServiceImp,
	)
	return config.ContainerServiceImp{}
}
