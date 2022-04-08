package addcontractinteraction

import (
	"fmt"

	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func AddContractInteractionToBlockChain(_blockChain []block.ContractInteractionsBlock, currentBlockId int, _contractInteraction contract.ContractInteraction) []block.ContractInteractionsBlock {

	_blockChain[currentBlockId].Data = append(_blockChain[currentBlockId].Data, _contractInteraction)

	fmt.Printf("Contract Interaction added -> %d\n", len(_blockChain[currentBlockId].Data))

	return _blockChain

}
