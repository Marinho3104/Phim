package addcontract

import "github.com/Marinho3104/Phim/src/structs/apiServer"

func SetRoute(serverPriv *apiServer.ApiServer) {

	serverPriv.Route.POST("/addContract", AddContractRouteFunction(serverPriv))

}
