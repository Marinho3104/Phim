package work

import (
	"fmt"
	"strings"

	"github.com/Marinho3104/Phim-BlockChain/src/comiteClient/Work/handleWorkFunctions/synchronization"
	"github.com/Marinho3104/Phim-BlockChain/src/structs/comite"
)

func HandleWork(comiteClient *comite.ComiteClient, workReceived []string) (all []byte) {

	all = []byte("")

	for index := 0; index < len(workReceived); index += 2 {
		if len(workReceived)-1 == index {
			all = []byte(workReceived[index])
			return
		}

		if workReceived[index] == "Start Synchronize" {
			fmt.Println("Sync received")

			index := synchronization.HandleSynchronization(comiteClient, workReceived, index)

			if index == -1 {
				all = []byte(strings.Join(workReceived[index:], "\n"))
				return
			}

		}
	}

	return
}
