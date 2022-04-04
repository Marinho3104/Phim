package maincomiteserverfunctions

import (
	"fmt"
	"net"
	"strings"

	"github.com/Marinho3104/Phim-BlockChain/src/structs/block"
	"github.com/Marinho3104/Phim-BlockChain/src/structs/comite"
	"github.com/Marinho3104/Phim-BlockChain/src/structs/transaction"
)

func Initialization(addr string, port string) *comite.ComiteServer {

	_listenner, err := createListenner(addr, port)

	if err != nil {
		panic(err)
	}

	comiteServer := &comite.ComiteServer{
		ComiteServerListener: _listenner,
		ComiteClientList:     make(chan []comite.ComiteClientConnection, 1),

		BlockChainTransaction: make(chan []block.TransactionBlock, 1),
		TransactionInCheck:    make(chan []transaction.ConfirmationTransaction, 1),
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

	comiteServer.ComiteClientList <- []comite.ComiteClientConnection{}

	comiteServer.BlockChainTransaction <- []block.TransactionBlock{}

}
