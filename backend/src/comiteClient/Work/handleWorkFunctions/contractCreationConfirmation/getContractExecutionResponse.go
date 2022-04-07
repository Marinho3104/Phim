package contractcreationconfirmation

import (
	contractfunctions "github.com/Marinho3104/Phim/src/comiteClient/commonComiteClient/contractFunctions"
	transactionblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/transactionBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func GetContractExecutionResponse(comiteClient *comite.ComiteClient, _contract contract.Contract) (map[string]interface{}, error) {

	_blockChain := <-comiteClient.BlockChainTransaction

	comiteClient.BlockChainTransaction <- _blockChain

	contractfunctions.WriteContract(_contract, transactionblockchain.GetBalance(_blockChain, _contract.ContractAddress), _contract.CreatorAddress, 0, _contract.Fee, make(map[string]interface{}), true)

	contractfunctions.WriteExecuter(_contract, "", make(map[string]interface{}))

	resp, err := contractfunctions.ExecuteContract()

	contractfunctions.DeleteContractsFile()

	if err != nil {
		return nil, err
	}

	response, err := contractfunctions.HandleContractResponse(resp)

	return response, err
}
