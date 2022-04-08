package sendcontractinteraction

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

func CreateConfirmationVariableAndSendToComite(comiteServer *comite.ComiteServer, _contractInteraction contract.ContractInteraction, _comiteList []comiteclientconnection.ComiteClientConnection) []byte {

	comiteChoose := sendtocomiteclient.GetComiteChoose(_comiteList, rand.Intn(len(_comiteList)))

	_contractInteraction.ComiteReviewerAddress = comiteChoose.Address

	_contractInteractionByteForm, _ := json.Marshal(_contractInteraction)

	hash := sha256.Sum256(_contractInteractionByteForm)

	_contractConfirmation := contract.ConfirmationContract{
		Hash256:        hex.EncodeToString(hash[:]),
		Interaction:    _contractInteraction,
		ComiteChoose:   comiteChoose,
		ComiteTotal:    _comiteList,
		ComiteResponse: []comiteclientconnection.ComiteClientConnection{},
		Resp:           make(map[string]int),
	}

	comiteServer.ContractConfirmation <- append(<-comiteServer.ContractConfirmation, _contractConfirmation)

	return bytes.Join(
		[][]byte{
			[]byte("Contract Interaction-" + _contractConfirmation.Hash256),
			_contractInteractionByteForm,
			[]byte(""),
		},
		[]byte("\n"),
	)
}
