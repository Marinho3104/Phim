package maincomiteserver

import (
	"fmt"
	"net"
	"strings"

	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func Initialization(addr, port string) (comiteServer *comite.ComiteServer) {

	_listenner, err := net.Listen("tcp", strings.Join(
		[]string{
			addr,
			port,
		},
		":"),
	)

	if err != nil {
		panic(err)
	}

	comiteServer = &comite.ComiteServer{
		ComiteServerListener:  _listenner,
		ComiteClientList:      make(chan []comiteclientconnection.ComiteClientConnection, 1),
		BlockChainTransaction: make(chan []block.TransactionBlock, 1),
	}

	initializaVariablesComiteServer(comiteServer)

	fmt.Println("Comite Server launch")

	return

}

func initializaVariablesComiteServer(comiteServer *comite.ComiteServer) {

	comiteServer.ComiteClientList <- []comiteclientconnection.ComiteClientConnection{}

	comiteServer.BlockChainTransaction <- []block.TransactionBlock{}

}
