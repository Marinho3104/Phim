package addcontractinteraction

import (
	"encoding/json"
	"fmt"
	"net/http"

	sendcontractinteraction "github.com/Marinho3104/Phim/src/comiteServer/sendToComiteClient/sendContractInteraction"
	"github.com/Marinho3104/Phim/src/structs/apiServer"
	"github.com/Marinho3104/Phim/src/structs/contract"
	"github.com/gin-gonic/gin"
)

func AddContractRouteFunction(_server *apiServer.ApiServer) gin.HandlerFunc {

	return func(c *gin.Context) {

		amount := -1

		fee := -1

		cValue := -1

		_arguments := map[string]interface{}{}

		fmt.Sscan(c.PostForm("amount"), &amount)

		fmt.Sscan(c.PostForm("c"), &cValue)

		fmt.Sscan(c.PostForm("fee"), &fee)

		err := json.Unmarshal([]byte(c.PostForm("args")), &_arguments)

		if err != nil {

			fmt.Println(err)
			return
		}

		var _currentContractInteraction contract.ContractInteraction = contract.ContractInteraction{

			ContractAddress: c.PostForm("contractAddress"),
			Interactor:      c.PostForm("address"),
			Amount:          amount,
			Function:        c.PostForm("function"),
			Arguments:       _arguments,
			Fee:             fee,
			Sign:            c.PostForm("sign"),
			C:               cValue,
		}

		sendcontractinteraction.GetContractInteraction(_server.Comite_Server, _currentContractInteraction)

		c.String(http.StatusOK, "Done")

	}

}
