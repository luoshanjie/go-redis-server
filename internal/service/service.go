package service

import "go-redis-server/internal/domain"

type baseService struct {
	tcp domain.TCPDelivery
}

func NewBaseService(tcp domain.TCPDelivery) *baseService {
	return &baseService{tcp: tcp}
}

// -----------------------------------------------------------------------------
//
// -----------------------------------------------------------------------------

func (that *baseService) Run() error {
	return that.tcp.ListenAndServer()
}

var _ domain.BaseService = new(baseService)
