package commoncomiteserver

import (
	"errors"

	"github.com/Marinho3104/Phim/src/common"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func RemoveComiteClient(comiteServer *comite.ComiteServer, _connClient comite.ComiteClientConnection) error {

	_newComiteList := []comite.ComiteClientConnection{}
	_removed := false

	for _, v := range <-comiteServer.ComiteClientList {

		if v.Connection == _connClient.Connection {
			_removed = true
			continue
		}

		_newComiteList = append(_newComiteList, v)
	}

	comiteServer.ComiteClientList <- _newComiteList

	return common.Trenary(_removed, nil, errors.New("CANNOT FIND THE PROVIDE CONNECTION IN COMITE LIST")).(error)

}
