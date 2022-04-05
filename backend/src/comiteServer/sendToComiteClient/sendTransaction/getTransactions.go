package sendtransaction

import (
	"sort"

	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func GetTransactions(comiteServer *comite.ComiteServer) {

	_blockChainTransactionTemp := <-comiteServer.BlockChainTransaction

	comiteServer.BlockChainTransaction <- _blockChainTransactionTemp

	_lastTransaction := _blockChainTransactionTemp[comiteServer.CurrentBlockId].Data[len(_blockChainTransactionTemp[comiteServer.CurrentBlockId].Data)-1]

	for {

		_transactions := <-comiteServer.TransactionInCheck

		sort.SliceStable(_transactions, func(p, q int) bool {
			return _transactions[p].Time < _transactions[q].Time
		})

		for _, v := range _transactions {

			if v.Time >= _lastTransaction.Time {

				_lastTransaction = v

				go CreateConfirmationVariableAndSendToComite(comiteServer, v)
			}

		}

		comiteServer.TransactionInCheck <- []transaction.Transaction{}

	}

}
