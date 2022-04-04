package confirmContractOwnershipafter

import (
	comiteserversynchronization "github.com/Marinho3104/Phim/src/comiteServer/requestReceived/comiteServerSynchronization"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func ContractOwnershipAccepted(comiteServer *comite.ComiteServer, _connClient comite.ComiteClientConnection) {
	comiteserversynchronization.ComiteServerSynchronization(comiteServer, _connClient)
}
