package delivery

import (
	"fmt"
	"go-redis-server/internal/domain"
	"go-redis-server/pkg/configure"
	"net"
)

type tcpDelivery struct {
	properties *configure.ServerProperties
}

func NewTCPDelivery(properties *configure.ServerProperties) domain.TCPDelivery {
	return &tcpDelivery{properties: properties}
}

// -----------------------------------------------------------------------------
//
// -----------------------------------------------------------------------------

func (that *tcpDelivery) ListenAndServeWitSignal() error {
	address := fmt.Sprintf("%s:%d", that.properties.Bind, that.properties.Port)
	print("listen: ", address, "\n")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	for {
		_, err := listener.Accept()
		if err != nil {
			break
		}
		go func() {

		}()
	}

	return nil
}
