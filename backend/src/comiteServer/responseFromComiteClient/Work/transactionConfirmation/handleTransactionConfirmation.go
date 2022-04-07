package transactionconfirmation

import (
	transactionconfirmationdone "github.com/Marinho3104/Phim/src/comiteServer/responseFromComiteClient/Work/transactionConfirmation/transactionConfirmationDone"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func HandleTransactionConfirmation(comiteServer *comite.ComiteServer,
	_connComiteClient comiteclientconnection.ComiteClientConnection,
	hash256, resp string,
) {

	_newTransactionInCheck := []transaction.ConfirmationTransaction{}

	_currentTransactionInCheck := <-comiteServer.TransactionInConfirmation

	_verified := false

	for _, v := range _currentTransactionInCheck {

		if v.Hash256 == hash256 && !_verified {
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

				_verified = true

				continue

			} else if float64(len(v.ComiteDeclined)) >= float64(len(v.ComiteTotal))*.5 && float64(len(v.ComiteDeclined)-1) < float64(len(v.ComiteTotal))*.5 {
				transactionconfirmationdone.TransactionDeclined(comiteServer, v)
				_verified = true
				continue

			}
		}

		_newTransactionInCheck = append(_newTransactionInCheck, v)
	}

	comiteServer.TransactionInConfirmation <- _newTransactionInCheck

}
