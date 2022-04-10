package contractautoexec

import (
	"encoding/json"
	"fmt"
	"reflect"

	handleworkfunctions "github.com/Marinho3104/Phim/src/comiteClient/Work/handleWorkFunctions"
	addcontractinteraction "github.com/Marinho3104/Phim/src/common/blockChainChanges/contract/addContractToBlockChain/addContractInteraction"
	contractblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/contractBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func HandleContractAutoExec(comiteClient *comite.ComiteClient, contractInteractionConfirmationRequest []string) {

	_contractInteraction := contract.ContractInteraction{}

	err := json.Unmarshal([]byte(contractInteractionConfirmationRequest[1]), &_contractInteraction)

	if err != nil {

		fmt.Println(err)

		return

	}

	_blockChainContractAutoExec := <-comiteClient.BlockChainContractAutoExec

	comiteClient.BlockChainContractAutoExec <- _blockChainContractAutoExec

	_contract, err := contractblockchain.GetContract(_blockChainContractAutoExec, _contractInteraction.ContractAddress)

	if err != nil {
		fmt.Println(err)
		return
	}

	_blockChainContractInteraction := <-comiteClient.BlockChainContractInteractions

	comiteClient.BlockChainContractInteractions <- _blockChainContractInteraction

	_lastInteraction, err := contractblockchain.GetLastContractInteraction(_blockChainContractInteraction, _contractInteraction.ContractAddress)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := GetExecutionResponse(comiteClient, _contract, _contractInteraction, _lastInteraction.Variables)

	comiteContractResponse := contract.ConfirmationContractComiteClient{
		Status:   false,
		Response: make(map[string]interface{}),
	}

	_var := make(map[string]interface{})

	for s, v := range resp {

		if s != "_PhimChainContract__operations" {
			_var[s] = v
		}
	}

	if err == nil && (!reflect.DeepEqual(_var, _lastInteraction.Variables) || !reflect.DeepEqual([]interface{}{}, resp["_PhimChainContract__operations"])) {

		_contractInteraction.Variables = _var

		comiteContractResponse.Status = true
		comiteContractResponse.Response = resp

		comiteClient.BlockChainContractInteractions <- addcontractinteraction.AddContractInteractionToBlockChain(<-comiteClient.BlockChainContractInteractions, comiteClient.CurrentBlockId, _contractInteraction)

	}

	comiteContractResponseByteForm, _ := json.Marshal(comiteContractResponse)

	handleworkfunctions.SendResponseToComiteServer(comiteClient.Connection, []string{contractInteractionConfirmationRequest[0], string(comiteContractResponseByteForm)})

}
