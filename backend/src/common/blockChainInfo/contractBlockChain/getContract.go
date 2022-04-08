package contractblockchain

import (
	"errors"

	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func GetContract(contractBlockChain []block.ContractBlock, address string) (contract.Contract, error) {

	for b := len(contractBlockChain) - 1; b >= 0; b-- {

		for t := len(contractBlockChain[b].Data) - 1; t >= 0; t-- {

			if contractBlockChain[b].Data[t].ContractAddress == address {
				return contractBlockChain[b].Data[t], nil
			}

		}

	}

	return contract.Contract{}, errors.New("CONTRACT ADDRESS NOT FOUND")

}
