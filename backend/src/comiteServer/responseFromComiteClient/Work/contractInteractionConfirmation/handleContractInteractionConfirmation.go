package contractinteractionconfirmation

import (
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func HandleContractInteractionConfirmation(comiteServer *comite.ComiteServer,
	_connComiteClient comiteclientconnection.ComiteClientConnection,
	hash256, resp string,
) {
	_newContractInCheck := []contract.ConfirmationContract{}

	_currentContractInCheck := <-comiteServer.ContractConfirmation

	_verified := false

	for _, v := range _currentContractInCheck {

		if v.Hash256 == hash256 && !_verified {

			if val, ok := v.Resp[resp]; ok {
				v.Resp[resp] = val + 1
			} else {
				v.Resp[resp] = 1
			}

			if float64(len(v.Resp)) >= float64(len(v.ComiteTotal))*.75 && float64(len(v.Resp)-1) < float64(len(v.ComiteTotal))*.75 {

				NecessaryResponsesReceived(comiteServer, v)

				continue

			}

		}

		_newContractInCheck = append(_newContractInCheck, v)
	}

	comiteServer.ContractConfirmation <- _newContractInCheck
}
