package comiteserversynchronization

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	commoncomiteserver "github.com/Marinho3104/Phim-BlockChain/src/comiteServer/commonComiteServer/removeVariable"
	"github.com/Marinho3104/Phim-BlockChain/src/structs/comite"
)

func ComiteServerSynchronization(comiteServer *comite.ComiteServer, _connClient comite.ComiteClientConnection) {

	msgStartSync := "Start Synchronize"

	msgEndSync := "End Synchronize"

	comiteServer.ComiteClientList <- append(<-comiteServer.ComiteClientList, _connClient)

	time.Sleep(5 * time.Second)

	_blockChainTransactions := <-comiteServer.BlockChainTransaction

	comiteServer.BlockChainTransaction <- _blockChainTransactions

	_blockChainTransactionByteForm, err := json.Marshal(_blockChainTransactions)

	if err != nil {

		commoncomiteserver.RemoveComiteClient(comiteServer, _connClient)

		return

	}

	_connClient.Connection.Write(bytes.Join(
		[][]byte{
			[]byte(msgStartSync),
			_blockChainTransactionByteForm,
			[]byte(msgEndSync),
		},
		[]byte("\n"),
	),
	)

	fmt.Println("All blockchains sended and comit is added to blokcchain !!")

}
