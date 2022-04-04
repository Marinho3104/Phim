package acceptcomiteclientrequest

import (
	"fmt"

	"github.com/Marinho3104/Phim/src/structs/comite"
)

func AcceptComiteClient(comiteServer *comite.ComiteServer) {

	fmt.Println("Getting comiteclient request")

	for {

		_conn, err := comiteServer.ComiteServerListener.Accept()

		if err != nil {
			continue
		}

		go HandleComiteClientRequest(comiteServer, _conn)
	}

}
