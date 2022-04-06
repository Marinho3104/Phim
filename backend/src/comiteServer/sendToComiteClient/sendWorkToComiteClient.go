package sendtocomiteclient

import (
	"fmt"

	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func SendWorkToComiteClient(connComiteClient comiteclientconnection.ComiteClientConnection) {

	for {

		_work := <-connComiteClient.WorkToDo

		connComiteClient.WorkToDo <- []byte("")

		if len(_work) > 150000 {

			fmt.Println("Too big")

		} else {
			connComiteClient.Connection.Write(
				_work,
			)
		}
	}
}
