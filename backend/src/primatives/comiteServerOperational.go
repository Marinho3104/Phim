package primatives

import (
	getblockchainfromfile "github.com/Marinho3104/Phim/src/comiteServer/commonComiteServer/getBlockChainFromFile"
	maincomiteserverfunctions "github.com/Marinho3104/Phim/src/comiteServer/mainComiteServerFunctions"
	acceptcomiteclientrequest "github.com/Marinho3104/Phim/src/comiteServer/requestReceived/acceptComiteClientRequest"
)

func ComiteServerOperational() {
	comiteServer := maincomiteserverfunctions.Initialization("", "54321")

	err := getblockchainfromfile.GetBlockChainFromFile(comiteServer)

	if err != nil {
		panic(err)
	}

	acceptcomiteclientrequest.AcceptComiteClient(comiteServer)
}
