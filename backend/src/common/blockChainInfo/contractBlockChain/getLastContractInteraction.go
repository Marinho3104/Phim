package contractblockchain

import (
	"errors"

	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func GetLastContractInteraction(contractInteractionBlockChain []block.ContractInteractionsBlock, address string) (contract.ContractInteraction, error) {

	for b := len(contractInteractionBlockChain) - 1; b >= 0; b-- {

		for t := len(contractInteractionBlockChain[b].Data) - 1; t >= 0; t-- {

			if contractInteractionBlockChain[b].Data[t].ContractAddress == address {
				return contractInteractionBlockChain[b].Data[t], nil
			}

		}

	}

	return contract.ContractInteraction{}, errors.New("CONTRACT ADDRESS NOT FOUND")
}
