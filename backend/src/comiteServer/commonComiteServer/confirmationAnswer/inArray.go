package confirmationanswer

import (
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func InArray(arr []comiteclientconnection.ComiteClientConnection, _conn comiteclientconnection.ComiteClientConnection) bool {

	for _, v := range arr {
		if v.Connection == _conn.Connection {
			return true
		}
	}

	return false

}
