package sendtocomiteclient

import (
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func GetComiteChoose(group []comiteclientconnection.ComiteClientConnection, num int) comiteclientconnection.ComiteClientConnection {

	c := 0

	for _, v := range group {
		if c == num {
			return v
		}
		c++
	}

	return comiteclientconnection.ComiteClientConnection{}

}
