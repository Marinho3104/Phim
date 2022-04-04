package reponsefromcomiteclient

import (
	"fmt"
	"strings"

	commoncomiteserver "github.com/Marinho3104/Phim-BlockChain/src/comiteServer/commonComiteServer/removeVariable"
	"github.com/Marinho3104/Phim-BlockChain/src/structs/comite"
)

func GetResponse(comiteServer *comite.ComiteServer, _connClient comite.ComiteClientConnection) {

	all := []byte("")

	for {

		resp := make([]byte, 150000)

		_len, err := _connClient.Connection.Read(resp)

		if err != nil {
			fmt.Println("Comite Disconnect ...")

			commoncomiteserver.RemoveComiteClient(comiteServer, _connClient)

			return

		}

		resp = resp[:_len]

		all = append(all, resp...)

		respSplit := strings.Split(string(all), "\n")

		HandleComiteResponse(comiteServer, _connClient, respSplit)

	}

}
