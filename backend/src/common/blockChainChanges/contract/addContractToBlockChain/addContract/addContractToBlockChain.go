package addcontract

import (
	"fmt"

	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func AddContractToBlockChain(_blockChain []block.ContractBlock, currentBlockId int, _contract contract.Contract) []block.ContractBlock {

	_blockChain[currentBlockId].Data = append(_blockChain[currentBlockId].Data, _contract)

	fmt.Printf("Contract added -> %d\n", len(_blockChain[currentBlockId].Data))

	return _blockChain

}
