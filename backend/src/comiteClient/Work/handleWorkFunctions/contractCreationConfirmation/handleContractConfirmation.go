package contractcreationconfirmation

import (
	"encoding/json"

	handleworkfunctions "github.com/Marinho3104/Phim/src/comiteClient/Work/handleWorkFunctions"
	addcontract "github.com/Marinho3104/Phim/src/common/blockChainChanges/contract/addContractToBlockChain/addContract"
	addcontractinteraction "github.com/Marinho3104/Phim/src/common/blockChainChanges/contract/addContractToBlockChain/addContractInteraction"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
	"github.com/google/go-cmp/cmp"
)

func HandleContractConfirmation(comiteClient *comite.ComiteClient, contractConfirmationRequest []string) {

	_contract := contract.Contract{}

	err := json.Unmarshal([]byte(contractConfirmationRequest[1]), &_contract)

	if err != nil {
		return
	}

	if ConfirmContract() {

		resp, err := GetContractExecutionResponse(comiteClient, _contract)

		comiteContractResponse := contract.ConfirmationContractComiteClient{
			Status:   false,
			Response: make(map[string]interface{}),
		}

		_contractInteraction := contract.ContractInteraction{}

		if err == nil {

			_var := make(map[string]interface{})

			for s, v := range resp {

				if s != "_PhimChainContract__operations" {
					_var[s] = v
				}
			}

			_contractInteraction = contract.ContractInteraction{
				ContractAddress: _contract.ContractAddress,
				Interactor:      _contract.CreatorAddress,
				Amount:          0,
				Function:        "initialization",
				Arguments:       make(map[string]interface{}),
				Fee:             _contract.Fee,
				Sign:            _contract.Sign,
				C:               _contract.C,
				Variables:       _var,
			}

			comiteContractResponse.Status = true
			comiteContractResponse.Response = resp
		} else {
			_contract.ContractAddress = ""
		}

		//Handle with error
		comiteContractResponseByteForm, _ := json.Marshal(comiteContractResponse)

		comiteClient.BlockChainContract <- addcontract.AddContractToBlockChain(<-comiteClient.BlockChainContract, comiteClient.CurrentBlockId, _contract)

		if !cmp.Equal(_contractInteraction, contract.ContractInteraction{}) {
			comiteClient.BlockChainContractInteractions <- addcontractinteraction.AddContractInteractionToBlockChain(<-comiteClient.BlockChainContractInteractions, comiteClient.CurrentBlockId, _contractInteraction)
		}
		handleworkfunctions.SendResponseToComiteServer(comiteClient.Connection, []string{contractConfirmationRequest[0], string(comiteContractResponseByteForm)})

	}

}
