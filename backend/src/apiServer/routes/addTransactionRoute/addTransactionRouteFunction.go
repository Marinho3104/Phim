package addtransactionroute

import (
	"fmt"
	"net/http"
	"time"

	addtransactiontotransactionsincheck "github.com/Marinho3104/Phim/src/comiteServer/requestReceived/addTransactionToTransactionsInCheck"
	"github.com/Marinho3104/Phim/src/structs/apiServer"
	"github.com/Marinho3104/Phim/src/structs/transaction"
	"github.com/gin-gonic/gin"
)

func AddTransactionRouteFunction(_server *apiServer.ApiServer) gin.HandlerFunc {

	return func(c *gin.Context) {

		amount := 0
		cValue := 0
		fee := 0

		fmt.Sscan(c.PostForm("amount"), &amount)

		fmt.Sscan(c.PostForm("c"), &cValue)

		fmt.Sscan(c.PostForm("fee"), &fee)

		var _currentTransation transaction.Transaction = transaction.Transaction{
			Time:                  time.Now().UnixNano(),
			AddressFrom:           c.PostForm("addressFrom"),
			AddressTo:             c.PostForm("addressTo"),
			Amount:                amount,
			C:                     cValue,
			Sign:                  c.PostForm("sign"),
			Fee:                   fee,
			ComiteReviewerAddress: "",
		}

		addtransactiontotransactionsincheck.AddTransactionToTransactionsInCheck(&_server.Comite_Server, _currentTransation)

		c.String(http.StatusOK, "Done")

	}
}
