//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"

	delivery "go-redis-server/internal/delivery"
	conf "go-redis-server/pkg/configure"
)

func NewService() *baseService {
	panic(wire.Build(commonSet))
}

var commonSet = wire.NewSet(
	conf.NewServerProperties,
	delivery.NewTCPDelivery,
	NewBaseService,
)
