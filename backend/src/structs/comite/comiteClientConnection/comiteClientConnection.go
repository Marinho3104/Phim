package comiteclientconnection

import "net"

type ComiteClientConnection struct {
	Connection    net.Conn
	Address       string
	WorkToDo      chan []byte
	AnswerToCheck chan string
}
