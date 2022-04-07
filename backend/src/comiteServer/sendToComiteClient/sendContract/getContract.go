package sendcontract

import (
	sendtocomiteclient "github.com/Marinho3104/Phim/src/comiteServer/sendToComiteClient"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func GetContracts(comiteServer *comite.ComiteServer, _contract contract.Contract) {
	if !ConfirmContract() {
		return
	}

	_comiteList := <-comiteServer.ComiteClientList

	comiteServer.ComiteClientList <- _comiteList

	info := CreateConfirmationVariableAndSendToComite(comiteServer, _contract, _comiteList)

	sendtocomiteclient.SetInfoComiteClient(
		comiteServer,
		_comiteList,
		info,
	)
}
