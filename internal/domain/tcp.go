package domain

type TCPDelivery interface {
	ListenAndServeWitSignal() error
}
