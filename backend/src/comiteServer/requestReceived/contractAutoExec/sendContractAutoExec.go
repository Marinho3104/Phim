package contractautoexec

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"

	sendtocomiteclient "github.com/Marinho3104/Phim/src/comiteServer/sendToComiteClient"
	contractblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/contractBlockChain"
	getfee "github.com/Marinho3104/Phim/src/common/blockChainInfo/getFee"
	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/comite"
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func SendContractAutoExec(comiteServer *comite.ComiteServer, _blockChain []block.ContractBlock, _comiteList []comiteclientconnection.ComiteClientConnection) {

	for _, v := range _blockChain {

		for _, vv := range v.Data {

			_blockchaininteraction := <-comiteServer.BlockChainContractInteractions

			comiteServer.BlockChainContractInteractions <- _blockchaininteraction

			_lastContractInteraction, err := contractblockchain.GetLastContractInteraction(_blockchaininteraction, vv.ContractAddress)

			if err != nil {
				fmt.Println(err)
				return
			}

			_contractInteraction := contract.ContractInteraction{
				ContractAddress: vv.ContractAddress,
				Interactor:      "ªAutoexecutionª",
				Amount:          0,
				Function:        "autoExecution",
				Fee:             getfee.GetFee(),
				Sign:            "ªAutoexecutionª",
				C:               _lastContractInteraction.C + 1,
			}

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

			comiteServer.ContractConfirmationAutoExec <- append(<-comiteServer.ContractConfirmationAutoExec, _contractConfirmation)

			if err != nil {
				fmt.Println(err)
				return
			}

			info := bytes.Join(
				[][]byte{
					[]byte("Contract AutoExec-" + _contractConfirmation.Hash256),
					[]byte(_contractInteractionByteForm),
					[]byte(""),
				},
				[]byte("\n"),
			)

			sendtocomiteclient.SetInfoComiteClient(
				comiteServer,
				_comiteList,
				info,
			)

		}
	}

}
