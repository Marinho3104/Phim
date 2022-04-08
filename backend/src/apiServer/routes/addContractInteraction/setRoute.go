package addcontractinteraction

import "github.com/Marinho3104/Phim/src/structs/apiServer"

func SetRoute(serverPriv *apiServer.ApiServer) {

	serverPriv.Route.POST("/addContractInteraction", AddContractRouteFunction(serverPriv))

}
