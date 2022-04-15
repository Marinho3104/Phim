package acceptcomiteclient

import (
	"fmt"
	"net"

	"github.com/Marinho3104/Phim/src/comiteServer/request/acceptComiteClient/synchronization"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func HandleComiteAccept(comiteServer *comite.ComiteServer, _conn net.Conn) {

	contractInfo := make([]byte, 150000)

	_len, err := _conn.Read(contractInfo)

	if err != nil {
		return
	}

	contractInfo = contractInfo[:_len]

	if !ConfirmContractOwnerShip() {
		return
	}

	fmt.Println("Contract accepted")

	_connClient := comiteclientconnection.ComiteClientConnection{
		Connection: _conn,
		Address:    string(contractInfo),
		WorkToDo:   make(chan []byte, 1),
	}

	synchronization.SynchronizationToComiteClient(comiteServer, _connClient)

}
