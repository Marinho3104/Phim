package comiteserversynchronization

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	commoncomiteserver "github.com/Marinho3104/Phim/src/comiteServer/commonComiteServer/removeVariable"
	reponsefromcomiteclient "github.com/Marinho3104/Phim/src/comiteServer/responseFromComiteClient"
	sendtocomiteclient "github.com/Marinho3104/Phim/src/comiteServer/sendToComiteClient"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

func ComiteServerSynchronization(comiteServer *comite.ComiteServer, _connClient comiteclientconnection.ComiteClientConnection) {

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
			[]byte(""),
		},
		[]byte("\n"),
	),
	)

	fmt.Println("All blockchains sended and comit is added to blokcchain !!")

	go reponsefromcomiteclient.GetResponse(comiteServer, _connClient)

	sendtocomiteclient.SendWorkToComiteClient(_connClient)

}
