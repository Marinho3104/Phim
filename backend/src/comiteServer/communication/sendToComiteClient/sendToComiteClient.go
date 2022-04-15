package sendtocomiteclient

import comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"

func SendToComiteClient(_connClient comiteclientconnection.ComiteClientConnection) {

	for {

		_connClient.Connection.Write(
			<-_connClient.WorkToDo,
		)

	}

}
