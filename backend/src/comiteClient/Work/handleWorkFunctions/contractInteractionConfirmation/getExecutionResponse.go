package contractinteractionconfirmation

import (
	"fmt"

	contractfunctions "github.com/Marinho3104/Phim/src/comiteClient/commonComiteClient/contractFunctions"
	addtransactiontoblockchain "github.com/Marinho3104/Phim/src/common/blockChainChanges/transaction/addTransactionToBlockChain"
	transactionblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/transactionBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func GetExecutionResponse(comiteClient *comite.ComiteClient, _contract contract.Contract, _interaction contract.ContractInteraction, _currentVariables map[string]interface{}) (map[string]interface{}, error) {
	_blockChain := <-comiteClient.BlockChainTransaction

	comiteClient.BlockChainTransaction <- _blockChain

	contractfunctions.WriteContract(_contract,
		transactionblockchain.GetBalance(_blockChain, _contract.ContractAddress),
		transactionblockchain.GetC(_blockChain, _contract.ContractAddress)+1,
		_interaction.Interactor,
		_interaction.Amount,
		_interaction.Fee,
		_currentVariables,
		false,
	)

	contractfunctions.WriteExecuter(_contract, _interaction.Function, _interaction.Arguments)

	resp, err := contractfunctions.ExecuteContract()

	contractfunctions.DeleteContractsFile()

	if err != nil {
		return nil, err
	}

	fmt.Println(string(resp))

	response, err := contractfunctions.HandleContractResponse(resp)

	response["_PhimChainContract__operations"] = contractfunctions.ConfirmOperations(comiteClient, response, _interaction.ContractAddress)

	for _, v := range response["_PhimChainContract__operations"].([]interface{}) {

		_transaction := v.(map[string]interface{})

		transaction := transaction.Transaction{
			AddressFrom:           _transaction["addressFrom"].(string),
			AddressTo:             _transaction["addressTo"].(string),
			Amount:                int(_transaction["amount"].(float64)),
			C:                     int(_transaction["c"].(float64)),
			Fee:                   int(_transaction["fee"].(float64)),
			ComiteReviewerAddress: "",
		}

		comiteClient.BlockChainTransaction <- addtransactiontoblockchain.AddTransactionToBlockChain(<-comiteClient.BlockChainTransaction, comiteClient.CurrentBlockId, transaction)

	}

	return response, err
}
