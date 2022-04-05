package work

import (
	"strings"

	"github.com/Marinho3104/Phim/src/comiteClient/Work/handleWorkFunctions/synchronization"
	transactionconfirmation "github.com/Marinho3104/Phim/src/comiteClient/Work/handleWorkFunctions/transactionConfirmation"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func HandleWork(comiteClient *comite.ComiteClient, workReceived []string) (all []byte) {

	all = []byte("")

	for index := 0; index < len(workReceived); index += 2 {
		if len(workReceived)-1 == index {
			all = []byte(workReceived[index])
			return
		}

		if workReceived[index] == "Start Synchronize" {

			index = synchronization.HandleSynchronization(comiteClient, workReceived, index)

			if index == -1 {
				all = []byte(strings.Join(workReceived[index:], "\n"))
				return
			}

		} else {

			_headerSplit := strings.Split(workReceived[index], "-")

			if _headerSplit[0] == "Transaction To Confirm" {
				transactionconfirmation.HandleTransactionConfirmation(comiteClient, workReceived[index:index+2])
			}
		}
	}

	return
}
