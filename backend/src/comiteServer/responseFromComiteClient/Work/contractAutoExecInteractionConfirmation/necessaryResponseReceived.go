package contractautoexecinteractionconfirmation

import (
	"encoding/json"
	"fmt"

	confirmationcontract "github.com/Marinho3104/Phim/src/comiteServer/commonComiteServer/confirmationContract"
	addcontractinteraction "github.com/Marinho3104/Phim/src/common/blockChainChanges/contract/addContractToBlockChain/addContractInteraction"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func autoExecInteraction(comiteSever *comite.ComiteServer, confirmationContract contract.ConfirmationContract) {
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
		fmt.Println("Refeused")
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
		ContractAddress:       confirmationContract.Interaction.ContractAddress,
		Interactor:            confirmationContract.Interaction.Interactor,
		Amount:                confirmationContract.Interaction.Amount,
		Function:              confirmationContract.Interaction.Function,
		Arguments:             confirmationContract.Interaction.Arguments,
		Fee:                   confirmationContract.Interaction.Fee,
		Sign:                  confirmationContract.Interaction.Sign,
		C:                     confirmationContract.Interaction.C,
		Variables:             _variables,
		ComiteReviewerAddress: confirmationContract.Interaction.ComiteReviewerAddress,
	}

	comiteSever.BlockChainContractInteractions <- addcontractinteraction.AddContractInteractionToBlockChain(<-comiteSever.BlockChainContractInteractions, comiteSever.CurrentBlockId, _contractInteraction)

	confirmationcontract.HandleOperations(comiteSever, _operations)
}

func NecessaryResponsesReceived(comiteSever *comite.ComiteServer, confirmationContract contract.ConfirmationContract) {

	autoExecInteraction(comiteSever, confirmationContract)

}
