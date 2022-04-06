package sendtocomiteclient

import (
	"bytes"
	"fmt"
	"strings"

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

		fmt.Println("Count -> ", len(strings.Split(string(newInfo), "\n")))

	}

	// _valueBytes, err := json.Marshal(_value)

	// if err != nil {
	// 	return
	// }

	// for _, v := range comiteList {

	// 	v.Connection.Write(
	// 		bytes.Join(
	// 			[][]byte{
	// 				[]byte(_header),
	// 				_valueBytes,
	// 				[]byte(""),
	// 			},
	// 			[]byte("\n"),
	// 		),
	// 	)

	// }

}
