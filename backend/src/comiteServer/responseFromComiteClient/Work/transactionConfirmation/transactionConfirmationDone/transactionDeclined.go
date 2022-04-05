package transactionconfirmationdone

import (
	addtransactiontoblockchain "github.com/Marinho3104/Phim/src/common/blockChainChanges/transaction/addTransactionToBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func TransactionDeclined(comiteServer *comite.ComiteServer, _transactionConfirmation transaction.ConfirmationTransaction) {

	_transactionDeclined := transaction.Transaction{
		Time:                  _transactionConfirmation.Transation.Time,
		AddressFrom:           _transactionConfirmation.Transation.AddressFrom,
		AddressTo:             "",
		Amount:                0,
		C:                     _transactionConfirmation.Transation.C,
		Sign:                  _transactionConfirmation.Transation.Sign,
		Fee:                   _transactionConfirmation.Transation.Fee,
		ComiteReviewerAddress: _transactionConfirmation.ComiteChoose.Address,
	}

	comiteServer.BlockChainTransaction <- addtransactiontoblockchain.AddTransactionToBlockChain(<-comiteServer.BlockChainTransaction, comiteServer.CurrentBlockId, _transactionDeclined)

}
