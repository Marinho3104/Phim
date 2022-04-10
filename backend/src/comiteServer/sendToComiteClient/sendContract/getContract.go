package sendcontract

import (
	"fmt"

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

	if _contract.AutoExec {
		_contract.ContractAddress = _contract.CreatorAddress + "ªAutoExecute" + fmt.Sprintf("%d", _contract.C)
	} else {
		_contract.ContractAddress = _contract.CreatorAddress + "ª" + fmt.Sprintf("%d", _contract.C)
	}

	info := CreateConfirmationVariableAndSendToComite(comiteServer, _contract, _comiteList)

	sendtocomiteclient.SetInfoComiteClient(
		comiteServer,
		_comiteList,
		info,
	)
}
