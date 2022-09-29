package domain

//go:generate mockery --name=TCPDelivery
type TCPDelivery interface {
	ListenAndServeWitSignal() error
}
