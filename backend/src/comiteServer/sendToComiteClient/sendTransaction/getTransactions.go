package sendtransaction

import (
	sendtocomiteclient "github.com/Marinho3104/Phim/src/comiteServer/sendToComiteClient"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func GetTransactions(comiteServer *comite.ComiteServer, _transaction transaction.Transaction) {

	if !ConfirmTransaction(comiteServer, _transaction) {
		return
	}

	_comiteList := <-comiteServer.ComiteClientList

	comiteServer.ComiteClientList <- _comiteList

	info := CreateConfirmationVariableAndSendToComite(comiteServer, _transaction, _comiteList)

	sendtocomiteclient.SetInfoComiteClient(
		comiteServer,
		_comiteList,
		info,
	)
}
