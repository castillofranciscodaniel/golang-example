//go:build InitializeServer
// +build InitializeServer

package main

import (
	service "bitbucket.org/chattigodev/chattigo-golang-library/pkg/deserializer"
	service2 "bitbucket.org/chattigodev/chattigo-golang-library/pkg/serializer"
	"github.com/castillofranciscodaniel/golang-example/config"
	"github.com/castillofranciscodaniel/golang-example/pkg/client"
	"github.com/castillofranciscodaniel/golang-example/pkg/handler"
	service3 "github.com/castillofranciscodaniel/golang-example/pkg/service"
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
)

func InitializeServer() config.ContainerService {
	wire.Build(
		service.NewDeserializerServiceImpl,
		service2.NewSerializerServiceImpl,
		resty.New,
		client.NewMessageChannelClientImpl,
		service3.NewMessageChannelServiceIml,
		handler.NewHealthHandler,
		handler.NewMessageChannelHandler,

		config.NewContainerServiceImp,
	)
	return config.ContainerService{}
}
