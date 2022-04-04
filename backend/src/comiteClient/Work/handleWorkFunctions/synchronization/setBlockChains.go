package synchronization

import (
	"encoding/json"

	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func SetBlockChains(comiteClient *comite.ComiteClient, _blockChainsInfo []string) error {

	_transactionBlockChain := []block.TransactionBlock{}

	err := json.Unmarshal([]byte(_blockChainsInfo[0]), &_transactionBlockChain)

	if err != nil {
		return err
	}

	comiteClient.CurrentBlockId = len(_transactionBlockChain) - 1

	comiteClient.BlockChainTransaction <- _transactionBlockChain

	comiteClient.Synchronized = true

	return nil

}
