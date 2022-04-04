package transaction

import "net"

type ConfirmationTransaction struct {
	Hash256              string
	Transation           Transaction
	ComiteChoose         net.Conn
	ComiteChooseResponse int
	ComiteChooseAddress  string
	ComiteTotal          []net.Conn
	ComiteAccepted       []net.Conn
	ComiteDeclined       []net.Conn
}
