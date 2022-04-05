package addTransactionToTransactionsInCheck

import (
	transactionblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/transactionBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func ConfirmTransaction(comiteServer *comite.ComiteServer, _transaction transaction.Transaction) bool {

	_blockChain := <-comiteServer.BlockChainTransaction

	comiteServer.BlockChainTransaction <- _blockChain

	cCheck := transactionblockchain.GetC(_blockChain, _transaction.AddressFrom)+1 == _transaction.C

	return cCheck
}
