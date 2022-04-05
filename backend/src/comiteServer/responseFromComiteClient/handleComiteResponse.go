package reponsefromcomiteclient

import (
	"strings"

	transactionconfirmation "github.com/Marinho3104/Phim/src/comiteServer/responseFromComiteClient/Work/transactionConfirmation"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func HandleComiteResponse(comiteServer *comite.ComiteServer, _connClient comiteclientconnection.ComiteClientConnection, _responseSplit []string) (all []byte) {

	all = []byte("")

	for index := 0; index < len(_responseSplit); index += 2 {

		if len(_responseSplit)-1 == index {
			all = []byte(_responseSplit[index])
			return
		}

		headerSplit := strings.Split(_responseSplit[index], "-")

		if headerSplit[0] == "Transaction To Confirm" {
			transactionconfirmation.HandleTransactionConfirmation(comiteServer, _connClient, headerSplit[1], _responseSplit[index+1])
		}
	}

	return
}
