package addTransactionToTransactionsInCheck

import (
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func AddTransactionToTransactionsInCheck(comiteServer *comite.ComiteServer, _transaction transaction.Transaction) {

	if ConfirmTransaction(comiteServer, _transaction) {
		comiteServer.TransactionInCheck <- append(<-comiteServer.TransactionInCheck, _transaction)
	}

}
