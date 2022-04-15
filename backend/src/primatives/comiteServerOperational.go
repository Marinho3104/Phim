package primatives

import (
	maincomiteserver "github.com/Marinho3104/Phim/src/comiteServer/mainComiteServer"
	acceptcomiteclient "github.com/Marinho3104/Phim/src/comiteServer/request/acceptComiteClient"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func ComiteServerOperational(chn chan *comite.ComiteServer) {
	comiteServer := maincomiteserver.Initialization("", "54321")

	chn <- comiteServer

	acceptcomiteclient.AcceptComiteClient(comiteServer)

}
