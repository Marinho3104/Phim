package work

import (
	"strings"

	contractconfirmation "github.com/Marinho3104/Phim/src/comiteClient/Work/handleWorkFunctions/contractCreationConfirmation"
	"github.com/Marinho3104/Phim/src/comiteClient/Work/handleWorkFunctions/synchronization"
	transactionconfirmation "github.com/Marinho3104/Phim/src/comiteClient/Work/handleWorkFunctions/transactionConfirmation"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func HandleWork(comiteClient *comite.ComiteClient) {

	for {

		_checker := true

		_workToDo := <-comiteClient.WorkInQueue

		respSplit := strings.Split(string(_workToDo), "\n")

		for index := 0; index < len(respSplit); index += 2 {

			if index == len(respSplit)-1 {
				comiteClient.WorkInQueue <- respSplit[index]
				_checker = false
				break
			}

			if respSplit[index] == "Start Synchronize" {

				index = synchronization.HandleSynchronization(comiteClient, respSplit, index)

				if index == -1 {
					comiteClient.WorkInQueue <- respSplit[index]
					_checker = false
					break
				}

			} else {

				_headerSplit := strings.Split(respSplit[index], "-")

				if _headerSplit[0] == "Transaction To Confirm" {

					transactionconfirmation.HandleTransactionConfirmation(comiteClient, respSplit[index:index+2])
				} else if _headerSplit[0] == "Contract Creation" {
					contractconfirmation.HandleContractConfirmation(comiteClient, respSplit[index:index+2])
				}
			}

		}

		if _checker {
			comiteClient.WorkInQueue <- ""
		}
	}

}
