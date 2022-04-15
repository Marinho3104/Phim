package comiteservercommunication

import "github.com/Marinho3104/Phim/src/structs/block"

type SynchronizationToComiteClient struct {
	Headers               map[string]interface{}
	TransactionBlockChain []block.TransactionBlock

	End bool
}
