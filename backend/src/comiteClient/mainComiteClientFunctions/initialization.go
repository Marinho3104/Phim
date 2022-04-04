package maincomiteclientfunctions

import (
	"net"
	"strings"

	"github.com/Marinho3104/Phim-BlockChain/src/structs/block"
	"github.com/Marinho3104/Phim-BlockChain/src/structs/comite"
)

func Initialization(addr, port string) *comite.ComiteClient {
	connection, err := createConnection(addr, port)

	if err != nil {
		panic(err)
	}

	comiteClient := &comite.ComiteClient{
		Connection:            connection,
		BlockChainTransaction: make(chan []block.TransactionBlock, 1),
		WorkAfterSynchronize:  make(chan []interface{}, 1),
	}

	initializaVariablesComiteClient(comiteClient)

	return comiteClient

}

func initializaVariablesComiteClient(comiteClient *comite.ComiteClient) {

	comiteClient.WorkAfterSynchronize <- []interface{}{}
}

func createConnection(addr string, port string) (net.Conn, error) {
	return net.Dial("tcp", strings.Join(
		[]string{
			addr,
			port,
		},
		":"),
	)
}
