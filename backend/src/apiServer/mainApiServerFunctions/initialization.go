package mainapiserverfunctions

import (
	comite "github.com/Marinho3104/Phim-BlockChain/src/structs/Comite"
	"github.com/Marinho3104/Phim-BlockChain/src/structs/apiServer"
	"github.com/gin-gonic/gin"
)

func Initialization(comiteServer *comite.ComiteServer) *apiServer.ApiServer {

	router := gin.Default()

	return &apiServer.ApiServer{
		Route:         router,
		Comite_Server: comiteServer,
	}

}
