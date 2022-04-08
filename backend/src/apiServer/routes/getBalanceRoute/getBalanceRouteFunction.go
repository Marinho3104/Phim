package getbalanceroute

import (
	"fmt"
	"net/http"

	contractblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/contractBlockChain"
	contractinteractionblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/contractInteractionBlockChain"
	transactionblockchain "github.com/Marinho3104/Phim/src/common/blockChainInfo/transactionBlockChain"
	"github.com/Marinho3104/Phim/src/structs/apiServer"
	"github.com/gin-gonic/gin"
)

func GetBalanceRouteFunction(_server *apiServer.ApiServer) gin.HandlerFunc {

	return func(c *gin.Context) {

		_blockChain := <-_server.Comite_Server.BlockChainTransaction

		_server.Comite_Server.BlockChainTransaction <- _blockChain

		_blockChainContract := <-_server.Comite_Server.BlockChainContract

		_server.Comite_Server.BlockChainContract <- _blockChainContract

		_blockChainContractInteraction := <-_server.Comite_Server.BlockChainContractInteractions

		_server.Comite_Server.BlockChainContractInteractions <- _blockChainContractInteraction

		_transactionBalance := transactionblockchain.GetBalance(_blockChain, c.PostForm("address"))

		_contractBalance := contractblockchain.GetBalance(_blockChainContract, c.PostForm("address"))

		_contractInteractionBalance := contractinteractionblockchain.GetBalance(_blockChainContractInteraction, c.PostForm("address"))

		c.String(http.StatusOK, fmt.Sprintf("%d", (_transactionBalance+_contractBalance+_contractInteractionBalance)))

	}

}
