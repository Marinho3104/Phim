package transactionblockchain

import "github.com/Marinho3104/Phim/src/structs/block"

func GetBalance(transactionBlockChain []block.TransactionBlock, address string) (balance int) {

	for b := len(transactionBlockChain) - 1; b >= 0; b-- {

		for t := len(transactionBlockChain[b].Data) - 1; t >= 0; t-- {

			if transactionBlockChain[b].Data[t].AddressFrom == address {

				balance -= transactionBlockChain[b].Data[t].Amount
				balance -= transactionBlockChain[b].Data[t].Fee
			}
			if transactionBlockChain[b].Data[t].AddressTo == address {
				balance += transactionBlockChain[b].Data[t].Amount
			}
			if transactionBlockChain[b].Data[t].ComiteReviewerAddress == address {
				balance += transactionBlockChain[b].Data[t].Fee
			}

		}

	}

	return

}
