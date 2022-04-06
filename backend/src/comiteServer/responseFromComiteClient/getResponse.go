package reponsefromcomiteclient

import (
	"fmt"
	"strings"

	commoncomiteserver "github.com/Marinho3104/Phim/src/comiteServer/commonComiteServer/removeVariable"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func GetResponse(comiteServer *comite.ComiteServer, _connClient comiteclientconnection.ComiteClientConnection) {

	all := []byte("")

	for {

		//fmt.Println("Getting response ...")

		resp := make([]byte, 15000000)

		_len, err := _connClient.Connection.Read(resp)

		if err != nil {
			fmt.Println("Comite Disconnect ...")

			commoncomiteserver.RemoveComiteClient(comiteServer, _connClient)

			return

		}

		//fmt.Println("Response obtained ...")

		resp = resp[:_len]

		all = append(all, resp...)

		respSplit := strings.Split(string(all), "\n")

		all = HandleComiteResponse(comiteServer, _connClient, respSplit)

	}

}
