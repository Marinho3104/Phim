package synchronization

import (
	"encoding/json"
	"fmt"

	sendtocomiteclient "github.com/Marinho3104/Phim/src/comiteServer/communication/sendToComiteClient"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
	comiteservercommunication "github.com/Marinho3104/Phim/src/structs/comite/comiteServerCommunication"
)

func SynchronizationToComiteClient(comiteServer *comite.ComiteServer, _connClient comiteclientconnection.ComiteClientConnection) {

	_header := map[string]interface{}{
		"content": "synchronization",
	}

	transactionBlockChain := <-comiteServer.BlockChainTransaction

	comiteServer.BlockChainTransaction <- transactionBlockChain

	_synchronization := comiteservercommunication.SynchronizationToComiteClient{
		Headers:               _header,
		TransactionBlockChain: transactionBlockChain,
		End:                   true,
	}

	_synchronizationByteForm, err := json.Marshal(_synchronization)

	if err != nil {
		fmt.Println(err)
		return
	}

	sendtocomiteclient.AddToComiteClientWork(&_connClient, _synchronizationByteForm)

	sendtocomiteclient.SendToComiteClient(_connClient)

}
