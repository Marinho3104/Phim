package contractautoexec

import (
	"fmt"
	"time"

	"github.com/Marinho3104/Phim/src/structs/comite"
)

func GetContractAutoExec(comiteServer *comite.ComiteServer) {

	for {

		_comiteList := <-comiteServer.ComiteClientList

		comiteServer.ComiteClientList <- _comiteList

		_blockChainContractAutoExec := <-comiteServer.BlockChainContractAutoExec

		comiteServer.BlockChainContractAutoExec <- _blockChainContractAutoExec

		if len(_comiteList) > 0 {

			SendContractAutoExec(comiteServer, _blockChainContractAutoExec, _comiteList)
		}

		fmt.Println("Wainting 10 seconds ...")

		time.Sleep(10 * time.Second)

	}

}
