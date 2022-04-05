package comite

import (
	"net"

	"github.com/Marinho3104/Phim/src/structs/block"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

type ComiteServer struct {
	ComiteServerListener net.Listener
	ComiteClientList     chan []comiteclientconnection.ComiteClientConnection

	TransactionGroup []comiteclientconnection.ComiteClientConnection

	CurrentBlockId int

	BlockChainTransaction     chan []block.TransactionBlock
	TransactionInCheck        chan []transaction.Transaction
	TransactionInConfirmation chan []transaction.ConfirmationTransaction
}
