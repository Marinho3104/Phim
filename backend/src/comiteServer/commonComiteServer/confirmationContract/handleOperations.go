package confirmationcontract

import (
	addtransactiontoblockchain "github.com/Marinho3104/Phim/src/common/blockChainChanges/transaction/addTransactionToBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func HandleOperations(comiteServer *comite.ComiteServer, operations map[string]interface{}) {

	for _, v := range operations["_PhimChainContract__operations"].([]interface{}) {
		addTransaction(comiteServer, v.(map[string]interface{}))
	}

}

func addTransaction(comiteServer *comite.ComiteServer, _transaction map[string]interface{}) {

	transaction := transaction.Transaction{
		AddressFrom:           _transaction["addressFrom"].(string),
		AddressTo:             _transaction["addressTo"].(string),
		Amount:                int(_transaction["amount"].(float64)),
		C:                     int(_transaction["c"].(float64)),
		Fee:                   int(_transaction["fee"].(float64)),
		ComiteReviewerAddress: "",
	}

	comiteServer.BlockChainTransaction <- addtransactiontoblockchain.AddTransactionToBlockChain(<-comiteServer.BlockChainTransaction, comiteServer.CurrentBlockId, transaction)

}
