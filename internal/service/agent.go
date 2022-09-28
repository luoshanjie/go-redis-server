package service

import "go-redis-server/internal/domain"

type agent struct {
	tcp domain.TCPDelivery
}

func NewAgent(tcp domain.TCPDelivery) domain.BaseService {
	return &agent{tcp: tcp}
}

// -----------------------------------------------------------------------------
//
// -----------------------------------------------------------------------------

func (that *agent) Run() error {
	return that.tcp.ListenAndServeWitSignal()
}
