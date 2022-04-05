package confirmContractOwnershipafter

import (
	comiteserversynchronization "github.com/Marinho3104/Phim/src/comiteServer/requestReceived/comiteServerSynchronization"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func ContractOwnershipAccepted(comiteServer *comite.ComiteServer, _connClient comiteclientconnection.ComiteClientConnection) {
	comiteserversynchronization.ComiteServerSynchronization(comiteServer, _connClient)
}
