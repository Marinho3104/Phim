package contractblockchain

import "github.com/Marinho3104/Phim/src/structs/block"

func GetBalance(contractBlockChain []block.ContractBlock, address string) (balance int) {

	balance = 0

	for b := len(contractBlockChain) - 1; b >= 0; b-- {

		for t := len(contractBlockChain[b].Data) - 1; t >= 0; t-- {

			if contractBlockChain[b].Data[t].CreatorAddress == address {
				balance -= contractBlockChain[b].Data[t].Fee

			}
			if contractBlockChain[b].Data[t].ComiteReviewerAddress == address {
				balance += contractBlockChain[b].Data[t].Fee
			}

		}

	}

	return
}
