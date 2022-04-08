package mainapiserverfunctions

import (
	addcontract "github.com/Marinho3104/Phim/src/apiServer/routes/addContract"
	addcontractinteraction "github.com/Marinho3104/Phim/src/apiServer/routes/addContractInteraction"
	addtransactionroute "github.com/Marinho3104/Phim/src/apiServer/routes/addTransactionRoute"
	getbalanceroute "github.com/Marinho3104/Phim/src/apiServer/routes/getBalanceRoute"
	"github.com/Marinho3104/Phim/src/structs/apiServer"
)

func AddRoutes(apiServer *apiServer.ApiServer) {

	addtransactionroute.SetRoute(apiServer)

	getbalanceroute.SetRoute(apiServer)

	addcontract.SetRoute(apiServer)

	addcontractinteraction.SetRoute(apiServer)
}
