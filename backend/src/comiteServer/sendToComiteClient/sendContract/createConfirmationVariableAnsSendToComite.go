package sendcontract

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/rand"

	sendtocomiteclient "github.com/Marinho3104/Phim/src/comiteServer/sendToComiteClient"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func CreateConfirmationVariableAndSendToComite(comiteServer *comite.ComiteServer, _contract contract.Contract, _comiteList []comiteclientconnection.ComiteClientConnection) []byte {

	comiteChoose := sendtocomiteclient.GetComiteChoose(_comiteList, rand.Intn(len(_comiteList)))

	_contract.ComiteReviewerAddress = comiteChoose.Address

	_contractByteForm, _ := json.Marshal(_contract)

	hash := sha256.Sum256(_contractByteForm)

	_contractConfirmation := contract.ConfirmationContract{
		Hash256:        hex.EncodeToString(hash[:]),
		Contract:       _contract,
		ComiteChoose:   comiteChoose,
		ComiteTotal:    _comiteList,
		ComiteResponse: []comiteclientconnection.ComiteClientConnection{},
		Resp:           make(map[string]int),
	}

	comiteServer.ContractConfirmation <- append(<-comiteServer.ContractConfirmation, _contractConfirmation)

	return bytes.Join(
		[][]byte{
			[]byte("Contract Creation-" + _contractConfirmation.Hash256),
			_contractByteForm,
			[]byte(""),
		},
		[]byte("\n"),
	)

}
