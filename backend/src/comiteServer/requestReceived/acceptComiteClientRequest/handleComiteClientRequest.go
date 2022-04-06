package acceptcomiteclientrequest

import (
	"net"

	confirmContractOwnershipafter "github.com/Marinho3104/Phim/src/comiteServer/requestReceived/acceptComiteClientRequest/confirmContractOwnershipAfter"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func HandleComiteClientRequest(comiteServer *comite.ComiteServer, _conn net.Conn) {

	contractInfo := make([]byte, 150000)

	_len, err := _conn.Read(contractInfo)

	if err != nil {
		return
	}

	contractInfo = contractInfo[:_len]

	if ConfirmContractOwnerShip(contractInfo) {

		_connClient := comiteclientconnection.ComiteClientConnection{
			Connection: _conn,
			Address:    string(contractInfo),
			WorkToDo:   make(chan []byte, 1),
		}

		_connClient.WorkToDo <- []byte("")

		confirmContractOwnershipafter.ContractOwnershipAccepted(comiteServer, _connClient)
	}

}
