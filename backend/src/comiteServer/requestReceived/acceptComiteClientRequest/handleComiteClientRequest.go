package acceptcomiteclientrequest

import (
	"net"

	confirmContractOwnershipafter "github.com/Marinho3104/Phim/src/comiteServer/requestReceived/acceptComiteClientRequest/confirmContractOwnershipAfter"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func HandleComiteClientRequest(comiteServer *comite.ComiteServer, _conn net.Conn) {

	contractInfo := make([]byte, 150000)

	_len, err := _conn.Read(contractInfo)

	if err != nil {
		return
	}

	contractInfo = contractInfo[:_len]

	if ConfirmContractOwnerShip(contractInfo) {

		_connClient := comite.ComiteClientConnection{
			Connection: _conn,
			Address:    string(contractInfo),
		}

		confirmContractOwnershipafter.ContractOwnershipAccepted(comiteServer, _connClient)
	}

}
