package sendtransaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/rand"

	sendtocomiteclient "github.com/Marinho3104/Phim/src/comiteServer/sendToComiteClient"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
	"github.com/Marinho3104/Phim/src/structs/transaction"
)

func CreateConfirmationVariableAndSendToComite(comiteServer *comite.ComiteServer, _transaction transaction.Transaction, _comiteList []comiteclientconnection.ComiteClientConnection) []byte {

	comiteChoose := sendtocomiteclient.GetComiteChoose(_comiteList, rand.Intn(len(_comiteList)))

	_transaction.ComiteReviewerAddress = comiteChoose.Address

	_transactionByteForm, _ := json.Marshal(_transaction)

	hash := sha256.Sum256(_transactionByteForm)

	_transactioConfirmation := transaction.ConfirmationTransaction{
		Hash256:        hex.EncodeToString(hash[:]),
		Transation:     _transaction,
		ComiteChoose:   comiteChoose,
		ComiteTotal:    _comiteList,
		ComiteAccepted: []comiteclientconnection.ComiteClientConnection{},
		ComiteDeclined: []comiteclientconnection.ComiteClientConnection{},
	}

	comiteServer.TransactionInConfirmation <- append(<-comiteServer.TransactionInConfirmation, _transactioConfirmation)

	return bytes.Join(
		[][]byte{
			[]byte("Transaction To Confirm-" + _transactioConfirmation.Hash256),
			_transactionByteForm,
			[]byte(""),
		},
		[]byte("\n"),
	)
}