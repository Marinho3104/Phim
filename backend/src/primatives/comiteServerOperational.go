package primatives

import (
	getblockchainfromfile "github.com/Marinho3104/Phim/src/comiteServer/commonComiteServer/getBlockChainFromFile"
	maincomiteserverfunctions "github.com/Marinho3104/Phim/src/comiteServer/mainComiteServerFunctions"
	acceptcomiteclientrequest "github.com/Marinho3104/Phim/src/comiteServer/requestReceived/acceptComiteClientRequest"
	contractautoexec "github.com/Marinho3104/Phim/src/comiteServer/requestReceived/contractAutoExec"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func ComiteServerOperational(chn chan *comite.ComiteServer) {
	comiteServer := maincomiteserverfunctions.Initialization("", "54321")

	chn <- comiteServer

	_len, err := getblockchainfromfile.GetBlockChainFromFile(comiteServer)

	if err != nil {
		panic(err)
	}

	comiteServer.CurrentBlockId = _len

	go acceptcomiteclientrequest.AcceptComiteClient(comiteServer)

	go contractautoexec.GetContractAutoExec(comiteServer)

	_chn := make(chan int, 1)

	<-_chn

}
