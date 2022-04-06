package comite

import (
	"net"

	"github.com/Marinho3104/Phim/src/structs/block"
)

type ComiteClient struct {
	Connection net.Conn

	Synchronized bool

	WorkInQueue chan string

	BlockChainTransaction chan []block.TransactionBlock
	WorkAfterSynchronize  chan []interface{}
	CurrentBlockId        int

	Count chan int
}
