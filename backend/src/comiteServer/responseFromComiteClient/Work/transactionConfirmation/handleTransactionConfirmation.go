package handleworkfunctions

import (
	confirmationanswer "github.com/Marinho3104/Phim/src/comiteServer/commonComiteServer/confirmationAnswer"
	transactionconfirmationdone "github.com/Marinho3104/Phim/src/comiteServer/responseFromComiteClient/Work/transactionConfirmation/transactionConfirmationDone"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func HandleTransactionConfirmation(comiteServer *comite.ComiteServer,
	_connComiteClient comiteclientconnection.ComiteClientConnection,
	hash256, resp string,
) {

	_newTransactionInConfirmationArr := []transaction.ConfirmationTransaction{}

	for _, v := range <-comiteServer.TransactionInConfirmation {

		if v.Hash256 == hash256 {

			if !confirmationanswer.InArray(v.ComiteTotal, _connComiteClient) || confirmationanswer.InArray(v.ComiteAccepted, _connComiteClient) || confirmationanswer.InArray(v.ComiteDeclined, _connComiteClient) {
				return
			}

			if resp == "Transaction Confirmed" {

				if v.ComiteChoose.Connection == _connComiteClient.Connection {
					v.ComiteChooseResponse = 1
				}

				v.ComiteAccepted = append(v.ComiteAccepted, _connComiteClient)

			} else {
				if v.ComiteChoose.Connection == _connComiteClient.Connection {
					v.ComiteChooseResponse = 0
				}

				v.ComiteDeclined = append(v.ComiteDeclined, _connComiteClient)
			}

			if float64(len(v.ComiteAccepted)) >= float64(len(v.ComiteTotal))*.5 && float64(len(v.ComiteAccepted)-1) < float64(len(v.ComiteTotal))*.5 {
				transactionconfirmationdone.TransactionAccepted(comiteServer, v)

				continue

			} else if float64(len(v.ComiteDeclined)) >= float64(len(v.ComiteTotal))*.5 && float64(len(v.ComiteDeclined)-1) < float64(len(v.ComiteTotal))*.5 {
				transactionconfirmationdone.TransactionDeclined(comiteServer, v)
				continue
			}

		}

		_newTransactionInConfirmationArr = append(_newTransactionInConfirmationArr, v)

	}

	comiteServer.TransactionInConfirmation <- _newTransactionInConfirmationArr

}
