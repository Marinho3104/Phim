package transactionblockchain

import "github.com/Marinho3104/Phim/src/structs/block"

func GetC(transactionBlockChain []block.TransactionBlock, address string) (c int) {

	c = -1

	for b := len(transactionBlockChain) - 1; b >= 0; b-- {

		for t := len(transactionBlockChain[b].Data) - 1; t >= 0; t-- {

			if transactionBlockChain[b].Data[t].AddressFrom == address {
				c = transactionBlockChain[b].Data[t].C
				return
			}

		}

	}

	return

}
