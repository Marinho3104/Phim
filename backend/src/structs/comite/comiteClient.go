package comite

import (
	"net"

	"github.com/Marinho3104/Phim/src/structs/block"
)

type ComiteClient struct {
	Connection net.Conn

	Synchronized bool

	WorkInQueue chan []byte

	BlockChainTransaction          chan []block.TransactionBlock
	BlockChainContract             chan []block.ContractBlock
	BlockChainContractAutoExec     chan []block.ContractBlock
	BlockChainContractInteractions chan []block.ContractInteractionsBlock
	WorkAfterSynchronize           chan []interface{}
	CurrentBlockId                 int

	Count chan int
}
