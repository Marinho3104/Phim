package contractinteractionconfirmation

import (
	"encoding/json"
	"fmt"

	"github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/Marinho3104/Phim/src/structs/contract"
)

func HandleContractInteractionCofirmation(comiteClient *comite.ComiteClient, contractInteractionConfirmationRequest []string) {
	fmt.Println("In")

	_contractInteraction := contract.ContractInteraction{}

	err := json.Unmarshal([]byte(contractInteractionConfirmationRequest[1]), &_contractInteraction)

	if err != nil {
		return
	}

	if ConfirmContractInteraction() {

	}
}
