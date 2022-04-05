package addtransactiontoblockchain

import (
	"fmt"

	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func AddTransactionToBlockChain(_blockChain []block.TransactionBlock, currentBlockId int, _transaction transaction.Transaction) []block.TransactionBlock {

	_blockChain[currentBlockId].Data = append(_blockChain[currentBlockId].Data, _transaction)

	fmt.Println(fmt.Sprintf("Transaction added -> %d", len(_blockChain[currentBlockId].Data)))

	return _blockChain

}
