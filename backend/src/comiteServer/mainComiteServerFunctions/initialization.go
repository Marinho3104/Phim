package maincomiteserverfunctions

import (
	"fmt"
	"net"
	"strings"

	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func Initialization(addr string, port string) *comite.ComiteServer {

	_listenner, err := createListenner(addr, port)

	if err != nil {
		panic(err)
	}

	comiteServer := &comite.ComiteServer{
		ComiteServerListener: _listenner,
		ComiteClientList:     make(chan []comiteclientconnection.ComiteClientConnection, 1),

		BlockChainTransaction:     make(chan []block.TransactionBlock, 1),
		TransactionInCheck:        make(chan []transaction.Transaction, 1),
		TransactionInConfirmation: make(chan []transaction.ConfirmationTransaction, 1),
		TransactionGroup:          []comiteclientconnection.ComiteClientConnection{},
	}

	initializaVariablesComiteServer(comiteServer)

	fmt.Println("Done")

	return comiteServer
}

func createListenner(addr string, port string) (net.Listener, error) {
	return net.Listen("tcp", strings.Join(
		[]string{
			addr,
			port,
		},
		":"),
	)
}

func initializaVariablesComiteServer(comiteServer *comite.ComiteServer) {

	comiteServer.ComiteClientList <- []comiteclientconnection.ComiteClientConnection{}

	comiteServer.BlockChainTransaction <- []block.TransactionBlock{}

	comiteServer.TransactionInCheck <- []transaction.Transaction{}

	comiteServer.TransactionInConfirmation <- []transaction.ConfirmationTransaction{}

}
