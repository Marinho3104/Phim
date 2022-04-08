package contractinteractionconfirmation

import (
	"encoding/json"

	confirmationcontract "github.com/Marinho3104/Phim/src/comiteServer/commonComiteServer/confirmationContract"
	addcontractinteraction "github.com/Marinho3104/Phim/src/common/blockChainChanges/contract/addContractToBlockChain/addContractInteraction"
	addtransactiontoblockchain "github.com/Marinho3104/Phim/src/common/blockChainChanges/transaction/addTransactionToBlockChain"
	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
	"github.com/Marinho3104/Phim/src/structs/transaction"
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

	_transaction := transaction.Transaction{
		AddressFrom:           confirmationContract.Interaction.Interactor,
		AddressTo:             confirmationContract.Interaction.ContractAddress,
		Amount:                confirmationContract.Interaction.Amount,
		C:                     -1,
		Sign:                  confirmationContract.Interaction.Sign,
		Fee:                   confirmationContract.Interaction.Fee,
		ComiteReviewerAddress: confirmationContract.Interaction.ComiteReviewerAddress,
	}

	comiteSever.BlockChainTransaction <- addtransactiontoblockchain.AddTransactionToBlockChain(<-comiteSever.BlockChainTransaction, comiteSever.CurrentBlockId, _transaction)

	if !comiteResponseMajority.Status {
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
