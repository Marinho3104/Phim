package maincomiteclient

import (
	"net"
	"strings"

	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func Initialization(addr, port string) (comiteClient *comite.ComiteClient) {

	_connection, err := net.Dial("tcp", strings.Join(
		[]string{
			addr,
			port,
		},
		":"),
	)

	if err != nil {
		panic(err)
	}

	comiteClient = &comite.ComiteClient{
		Connection:            _connection,
		WorkInQueue:           make(chan []byte, 1),
		Synchronized:          false,
		BlockChainTransaction: make(chan []block.TransactionBlock, 1),
	}

	initializaVariablesComiteClient(comiteClient)

	return

}

func initializaVariablesComiteClient(comiteClient *comite.ComiteClient) {

}
