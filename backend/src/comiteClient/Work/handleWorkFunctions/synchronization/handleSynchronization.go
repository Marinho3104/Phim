package synchronization

import (
	"fmt"

	"github.com/Marinho3104/Phim-BlockChain/src/structs/comite"
)

func HandleSynchronization(comiteClient *comite.ComiteClient, workReceived []string, index int) int {

	_endSyncIndex := -1

	for i := index + 1; i < len(workReceived); i++ {

		if workReceived[i] == "End Synchronize" {
			_endSyncIndex = i
			break
		}
	}

	if _endSyncIndex == -1 {
		return _endSyncIndex
	}

	err := SetBlockChains(comiteClient, workReceived[index+1:_endSyncIndex])

	if err != nil {
		panic("ERROR ON RECEIVED BLOCKCHAIN INFO")
	}

	fmt.Println("BlockChain Updated")

	return _endSyncIndex - 1

}
