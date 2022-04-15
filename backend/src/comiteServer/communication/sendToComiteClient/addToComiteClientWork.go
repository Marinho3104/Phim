package sendtocomiteclient

import (
	"bytes"

	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func AddToComiteClientWork(_connClient *comiteclientconnection.ComiteClientConnection, _info []byte) {

	select {

	case _currentWorkToDo := <-_connClient.WorkToDo:

		_connClient.WorkToDo <- bytes.Join(
			[][]byte{
				_currentWorkToDo,
				_info,
			},
			[]byte("\n"),
		)

	default:

		_connClient.WorkToDo <- _info

	}

}
