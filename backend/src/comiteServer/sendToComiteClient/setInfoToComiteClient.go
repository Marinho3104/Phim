package sendtocomiteclient

import (
	"bytes"

	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func SetInfoComiteClient(comiteServer *comite.ComiteServer, connComiteClient []comiteclientconnection.ComiteClientConnection, info []byte) {

	for _, v := range connComiteClient {
		newInfo := bytes.Join(
			[][]byte{
				<-v.WorkToDo,
				info,
			},
			[]byte(""),
		)

		v.WorkToDo <- newInfo

	}

}
