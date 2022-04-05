package sendtocomiteclient

import (
	"bytes"
	"encoding/json"

	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func SendToComiteClient(comiteServer *comite.ComiteServer, comiteList []comiteclientconnection.ComiteClientConnection, _header string, _value interface{}) {

	_valueBytes, err := json.Marshal(_value)

	if err != nil {
		return
	}

	for _, v := range comiteList {

		v.Connection.Write(
			bytes.Join(
				[][]byte{
					[]byte(_header),
					_valueBytes,
					[]byte(""),
				},
				[]byte("\n"),
			),
		)

	}

}
