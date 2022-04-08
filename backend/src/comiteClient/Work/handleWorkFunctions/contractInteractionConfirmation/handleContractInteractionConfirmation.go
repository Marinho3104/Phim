package contractinteractionconfirmation

import (
	"encoding/json"
	"fmt"

	handleworkfunctions "github.com/Marinho3104/Phim/src/comiteClient/Work/handleWorkFunctions"
	addcontractinteraction "github.com/Marinho3104/Phim/src/common/blockChainChanges/contract/addContractToBlockChain/addContractInteraction"
	addtransactiontoblockchain "github.com/Marinho3104/Phim/src/common/blockChainChanges/transaction/addTransactionToBlockChain"
	contractblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/contractBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func HandleContractInteractionCofirmation(comiteClient *comite.ComiteClient, contractInteractionConfirmationRequest []string) {

	_contractInteraction := contract.ContractInteraction{}

	err := json.Unmarshal([]byte(contractInteractionConfirmationRequest[1]), &_contractInteraction)

	if err != nil {
		return
	}

	if ConfirmContractInteraction() {

		_transaction := transaction.Transaction{
			AddressFrom:           _contractInteraction.Interactor,
			AddressTo:             _contractInteraction.ContractAddress,
			Amount:                _contractInteraction.Amount,
			C:                     -1,
			Sign:                  _contractInteraction.Sign,
			Fee:                   _contractInteraction.Fee,
			ComiteReviewerAddress: _contractInteraction.ComiteReviewerAddress,
		}

		comiteClient.BlockChainTransaction <- addtransactiontoblockchain.AddTransactionToBlockChain(<-comiteClient.BlockChainTransaction, comiteClient.CurrentBlockId, _transaction)

		_blockChainContract := <-comiteClient.BlockChainContract

		comiteClient.BlockChainContract <- _blockChainContract

		_blockChainContractInteraction := <-comiteClient.BlockChainContractInteractions

		comiteClient.BlockChainContractInteractions <- _blockChainContractInteraction

		_contract, err := contractblockchain.GetContract(_blockChainContract, _contractInteraction.ContractAddress)

		if err != nil {
			fmt.Println(err)
			return
		}

		_lastInteraction, err := contractblockchain.GetLastContractInteraction(_blockChainContractInteraction, _contractInteraction.ContractAddress)

		if err != nil {
			fmt.Println(err)
			return
		}

		resp, err := GetExecutionResponse(comiteClient, _contract, _contractInteraction, _lastInteraction.Variables)

		fmt.Println(resp)

		comiteContractResponse := contract.ConfirmationContractComiteClient{
			Status:   false,
			Response: make(map[string]interface{}),
		}

		if err == nil {

			_var := make(map[string]interface{})

			for s, v := range resp {

				if s != "_PhimChainContract__operations" {
					_var[s] = v
				}
			}

			_contractInteraction.Variables = _var

			comiteContractResponse.Status = true
			comiteContractResponse.Response = resp

			comiteClient.BlockChainContractInteractions <- addcontractinteraction.AddContractInteractionToBlockChain(<-comiteClient.BlockChainContractInteractions, comiteClient.CurrentBlockId, _contractInteraction)

		} else {
			fmt.Println(err)
		}

		comiteContractResponseByteForm, _ := json.Marshal(comiteContractResponse)

		handleworkfunctions.SendResponseToComiteServer(comiteClient.Connection, []string{contractInteractionConfirmationRequest[0], string(comiteContractResponseByteForm)})

	}
}
