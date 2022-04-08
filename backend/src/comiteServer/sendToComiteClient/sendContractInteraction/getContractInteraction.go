package sendcontractinteraction

import (
	sendtocomiteclient "github.com/Marinho3104/Phim/src/comiteServer/sendToComiteClient"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func GetContractInteraction(comiteServer *comite.ComiteServer, _contractInteraction contract.ContractInteraction) {

	if !ConfirmContractInteraction() {
		return
	}

	_comiteList := <-comiteServer.ComiteClientList

	comiteServer.ComiteClientList <- _comiteList

	info := CreateConfirmationVariableAndSendToComite(comiteServer, _contractInteraction, _comiteList)

	sendtocomiteclient.SetInfoComiteClient(
		comiteServer,
		_comiteList,
		info,
	)

}
