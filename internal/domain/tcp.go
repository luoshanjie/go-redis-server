package domain

//go:generate mockery --name=TCPDelivery
type TCPDelivery interface {
	ListenAndServer() error
}
