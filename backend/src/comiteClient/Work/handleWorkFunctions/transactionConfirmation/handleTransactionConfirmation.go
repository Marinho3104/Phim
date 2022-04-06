package transactionconfirmation

import (
	"encoding/json"
	"fmt"

	handleworkfunctions "github.com/Marinho3104/Phim/src/comiteClient/Work/handleWorkFunctions"
	addtransactiontoblockchain "github.com/Marinho3104/Phim/src/common/blockChainChanges/transaction/addTransactionToBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func HandleTransactionConfirmation(comiteClient *comite.ComiteClient, transactionConfirmationRequest []string) {

	_transaction := transaction.Transaction{}

	err := json.Unmarshal([]byte(transactionConfirmationRequest[1]), &_transaction)

	if err != nil {
		return
	}

	if !comiteClient.Synchronized {
		fmt.Println("In synchronization")
	} else {

		resp := ""

		if ConfirmTransaction(comiteClient, _transaction) {

			resp = "Transaction Confirmed"

		} else {

			_transaction = transaction.Transaction{
				AddressFrom:           _transaction.AddressFrom,
				AddressTo:             "",
				Amount:                0,
				C:                     _transaction.C,
				Sign:                  _transaction.Sign,
				Fee:                   _transaction.Fee,
				ComiteReviewerAddress: _transaction.ComiteReviewerAddress,
			}

			resp = "Transaction Not Confirmed"
		}

		comiteClient.BlockChainTransaction <- addtransactiontoblockchain.AddTransactionToBlockChain(<-comiteClient.BlockChainTransaction, comiteClient.CurrentBlockId, _transaction)

		handleworkfunctions.SendResponseToComiteServer(comiteClient.Connection, []string{transactionConfirmationRequest[0], resp})
	}

}
