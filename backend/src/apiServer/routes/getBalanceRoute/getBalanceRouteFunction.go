package getbalanceroute

import (
	"fmt"
	"net/http"

	transactionblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/transactionBlockChain"
	"github.com/Marinho3104/Phim/src/structs/apiServer"
	"github.com/gin-gonic/gin"
)

func GetBalanceRouteFunction(_server *apiServer.ApiServer) gin.HandlerFunc {

	return func(c *gin.Context) {

		_blockChain := <-_server.Comite_Server.BlockChainTransaction

		_server.Comite_Server.BlockChainTransaction <- _blockChain

		c.String(http.StatusOK, fmt.Sprintf("%d", transactionblockchain.GetBalance(_blockChain, c.PostForm("address"))))

	}

}
