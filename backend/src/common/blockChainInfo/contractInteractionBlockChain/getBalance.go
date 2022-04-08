package contractinteractionblockchain

import "github.com/Marinho3104/Phim/src/structs/block"

func GetBalance(contractInteractionBlockChain []block.ContractInteractionsBlock, address string) (balance int) {
	balance = 0

	for b := len(contractInteractionBlockChain) - 1; b >= 0; b-- {

		for t := len(contractInteractionBlockChain[b].Data) - 1; t >= 0; t-- {

			if contractInteractionBlockChain[b].Data[t].Interactor == address {
				balance -= contractInteractionBlockChain[b].Data[t].Fee

			}
			if contractInteractionBlockChain[b].Data[t].ComiteReviewerAddress == address {
				balance += contractInteractionBlockChain[b].Data[t].Fee
			}

		}

	}

	return
}
