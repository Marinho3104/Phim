package comite

import (
	"net"

	"github.com/Marinho3104/Phim-BlockChain/src/structs/block"
	"github.com/Marinho3104/Phim-BlockChain/src/structs/transaction"
)

type ComiteServer struct {
	ComiteServerListener net.Listener
	ComiteClientList     chan []ComiteClientConnection

	BlockChainTransaction chan []block.TransactionBlock
	TransactionInCheck    chan []transaction.ConfirmationTransaction
}
