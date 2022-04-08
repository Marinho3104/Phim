package reponsefromcomiteclient

import (
	"strings"

	contractcreationconfirmation "github.com/Marinho3104/Phim/src/comiteServer/responseFromComiteClient/Work/contractCreationConfirmation"
	contractinteractionconfirmation "github.com/Marinho3104/Phim/src/comiteServer/responseFromComiteClient/Work/contractInteractionConfirmation"
	transactionconfirmation "github.com/Marinho3104/Phim/src/comiteServer/responseFromComiteClient/Work/transactionConfirmation"

	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func HandleComiteResponse(comiteServer *comite.ComiteServer, _connClient comiteclientconnection.ComiteClientConnection) {

	for {
		_checker := true

		_answersToCheck := <-_connClient.AnswerToCheck

		_responseSplit := strings.Split(_answersToCheck, "\n")

		for index := 0; index < len(_responseSplit); index += 2 {

			if len(_responseSplit)-1 == index {

				_connClient.AnswerToCheck <- _responseSplit[index]
				_checker = false
				break

			}

			headerSplit := strings.Split(_responseSplit[index], "-")

			if headerSplit[0] == "Transaction To Confirm" {
				transactionconfirmation.HandleTransactionConfirmation(comiteServer, _connClient, headerSplit[1], _responseSplit[index+1])
			} else if headerSplit[0] == "Contract Creation" {
				contractcreationconfirmation.HandleContractCreationConfirmation(comiteServer, _connClient, headerSplit[1], _responseSplit[index+1])
			} else if headerSplit[0] == "Contract Interaction" {
				contractinteractionconfirmation.HandleContractInteractionConfirmation(comiteServer, _connClient, headerSplit[1], _responseSplit[index+1])
			}
		}

		if _checker {
			_connClient.AnswerToCheck <- ""
		}
	}

}
