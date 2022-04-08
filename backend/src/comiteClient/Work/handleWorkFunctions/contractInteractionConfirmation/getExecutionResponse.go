package contractinteractionconfirmation

import (
	"fmt"

	contractfunctions "github.com/Marinho3104/Phim/src/comiteClient/commonComiteClient/contractFunctions"
	transactionblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/transactionBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
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

	return response, err
}
