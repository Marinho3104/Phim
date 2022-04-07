package reponsefromcomiteclient

import (
	"fmt"

	commoncomiteserver "github.com/Marinho3104/Phim/src/comiteServer/commonComiteServer/removeVariable"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func GetResponse(comiteServer *comite.ComiteServer, _connClient comiteclientconnection.ComiteClientConnection) {

	for {

		//fmt.Println("Getting response ...")

		resp := make([]byte, 150000)

		_len, err := _connClient.Connection.Read(resp)

		if err != nil {
			fmt.Println("Comite Disconnect ...")

			commoncomiteserver.RemoveComiteClient(comiteServer, _connClient)

			return

		}

		resp = resp[:_len]

		_connClient.AnswerToCheck <- <-_connClient.AnswerToCheck + string(resp)

	}

}
