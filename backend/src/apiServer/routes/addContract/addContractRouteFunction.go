package addcontract

import (
	"fmt"
	"net/http"
	"strconv"

	sendcontract "github.com/Marinho3104/Phim/src/comiteServer/sendToComiteClient/sendContract"
	"github.com/Marinho3104/Phim/src/structs/apiServer"
	"github.com/Marinho3104/Phim/src/structs/contract"
	"github.com/gin-gonic/gin"
)

func AddContractRouteFunction(_server *apiServer.ApiServer) gin.HandlerFunc {

	return func(c *gin.Context) {

		cValue := -1
		fee := -1

		fmt.Sscan(c.PostForm("c"), &cValue)
		fmt.Sscan(c.PostForm("fee"), &fee)

		boolValue, err := strconv.ParseBool(c.PostForm("fee"))

		if err != nil {
			c.String(http.StatusOK, "Done")
			return
		}

		var _currentContract contract.Contract = contract.Contract{

			CreatorAddress:        c.PostForm("address"),
			ContractAddress:       "",
			AutoExec:              boolValue,
			ContractName:          c.PostForm("contractName"),
			Data:                  []byte(c.PostForm("data")),
			C:                     cValue,
			Sign:                  c.PostForm("sign"),
			Fee:                   fee,
			ComiteReviewerAddress: "",
		}

		sendcontract.GetContracts(_server.Comite_Server, _currentContract)

		c.String(http.StatusOK, "Done")

	}

}
