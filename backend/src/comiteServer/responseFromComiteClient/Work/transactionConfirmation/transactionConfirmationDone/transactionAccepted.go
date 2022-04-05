package transactionconfirmationdone

import (
	addtransactiontoblockchain "github.com/Marinho3104/Phim/src/common/blockChainChanges/transaction/addTransactionToBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func TransactionAccepted(comiteServer *comite.ComiteServer, _transactionConfirmation transaction.ConfirmationTransaction) {

	comiteServer.BlockChainTransaction <- addtransactiontoblockchain.AddTransactionToBlockChain(<-comiteServer.BlockChainTransaction, comiteServer.CurrentBlockId, _transactionConfirmation.Transation)

}
