// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/google/wire"
	"go-redis-server/internal/delivery"
	"go-redis-server/pkg/configure"
)

// Injectors from wire.go:

func NewService() *baseService {
	serverProperties := configure.NewServerProperties()
	tcpDelivery := delivery.NewTCPDelivery(serverProperties)
	serviceBaseService := NewBaseService(tcpDelivery)
	return serviceBaseService
}

// wire.go:

var commonSet = wire.NewSet(configure.NewServerProperties, delivery.NewTCPDelivery, NewBaseService)