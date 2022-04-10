package contractcreationconfirmation

import (
	"encoding/json"
	"fmt"

	confirmationcontract "github.com/Marinho3104/Phim/src/comiteServer/commonComiteServer/confirmationContract"
	addcontract "github.com/Marinho3104/Phim/src/common/blockChainChanges/contract/addContractToBlockChain/addContract"
	addcontractinteraction "github.com/Marinho3104/Phim/src/common/blockChainChanges/contract/addContractToBlockChain/addContractInteraction"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func NecessaryResponsesReceived(comiteSever *comite.ComiteServer, confirmationContract contract.ConfirmationContract) {

	comiteResponseMajorityString := ""
	_count := -1

	for s, v := range confirmationContract.Resp {

		if v > _count {
			comiteResponseMajorityString = s
			_count = v
		}

	}

	comiteResponseMajority := contract.ConfirmationContractComiteClient{}

	err := json.Unmarshal([]byte(comiteResponseMajorityString), &comiteResponseMajority)

	if err != nil {
		return
	}

	if !comiteResponseMajority.Status {

		confirmationContract.Contract.ContractAddress = ""

		if confirmationContract.Contract.AutoExec {
			comiteSever.BlockChainContractAutoExec <- addcontract.AddContractToBlockChain(<-comiteSever.BlockChainContractAutoExec, comiteSever.CurrentBlockId, confirmationContract.Contract)
		} else {
			comiteSever.BlockChainContract <- addcontract.AddContractToBlockChain(<-comiteSever.BlockChainContract, comiteSever.CurrentBlockId, confirmationContract.Contract)
		}

		return

	}

	_variables := make(map[string]interface{})

	_operations := make(map[string]interface{})

	for s, v := range comiteResponseMajority.Response {

		if s == "_PhimChainContract__operations" {
			_operations[s] = v
			continue
		}

		_variables[s] = v

	}

	_contractInteraction := contract.ContractInteraction{
		ContractAddress:       confirmationContract.Contract.ContractAddress,
		Interactor:            confirmationContract.Contract.CreatorAddress,
		Amount:                0,
		Function:              "initialization",
		Arguments:             make(map[string]interface{}),
		Fee:                   confirmationContract.Contract.Fee,
		Sign:                  confirmationContract.Contract.Sign,
		C:                     confirmationContract.Contract.C,
		Variables:             _variables,
		ComiteReviewerAddress: confirmationContract.ComiteChoose.Address,
	}

	if confirmationContract.Contract.AutoExec {
		fmt.Println("Auto exec")
		comiteSever.BlockChainContractAutoExec <- addcontract.AddContractToBlockChain(<-comiteSever.BlockChainContractAutoExec, comiteSever.CurrentBlockId, confirmationContract.Contract)
	} else {
		comiteSever.BlockChainContract <- addcontract.AddContractToBlockChain(<-comiteSever.BlockChainContract, comiteSever.CurrentBlockId, confirmationContract.Contract)
	}
	comiteSever.BlockChainContractInteractions <- addcontractinteraction.AddContractInteractionToBlockChain(<-comiteSever.BlockChainContractInteractions, comiteSever.CurrentBlockId, _contractInteraction)

	confirmationcontract.HandleOperations(comiteSever, _operations)
}
