package comite

import "net"

type ComiteClientConnection struct {
	Connection net.Conn
	Address    string
}
