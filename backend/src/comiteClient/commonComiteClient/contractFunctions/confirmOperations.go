package contractfunctions

import (
	contractblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/contractBlockChain"
	contractinteractionblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/contractInteractionBlockChain"
	transactionblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/transactionBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func ConfirmOperations(comiteClient *comite.ComiteClient, operations map[string]interface{}, contractAddress string) []interface{} {

	_operationsConfirmed := make([]interface{}, 0)

	_blockChain := <-comiteClient.BlockChainTransaction

	comiteClient.BlockChainTransaction <- _blockChain

	_blockChainContract := <-comiteClient.BlockChainContract

	comiteClient.BlockChainContract <- _blockChainContract

	_blockChainContractInteraction := <-comiteClient.BlockChainContractInteractions

	comiteClient.BlockChainContractInteractions <- _blockChainContractInteraction

	_transactionBalance := transactionblockchain.GetBalance(_blockChain, contractAddress)

	_contractBalance := contractblockchain.GetBalance(_blockChainContract, contractAddress)

	_contractInteractionBalance := contractinteractionblockchain.GetBalance(_blockChainContractInteraction, contractAddress)

	for _, v := range operations["_PhimChainContract__operations"].([]interface{}) {

		_map := v.(map[string]interface{})

		if (_transactionBalance+_contractBalance+_contractInteractionBalance) > int(_map["amount"].(float64)) && _map["addressFrom"] == contractAddress {
			_operationsConfirmed = append(_operationsConfirmed, v)
		}
	}

	return _operationsConfirmed

}
