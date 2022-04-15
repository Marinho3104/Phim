package acceptcomiteclient

import "github.com/Marinho3104/Phim/src/structs/comite"

func AcceptComiteClient(comiteServer *comite.ComiteServer) {

	for {

		_conn, err := comiteServer.ComiteServerListener.Accept()

		if err != nil {
			continue
		}

		HandleComiteAccept(comiteServer, _conn)

	}

}
