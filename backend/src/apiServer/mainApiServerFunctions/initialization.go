package mainapiserverfunctions

import (
	"github.com/Marinho3104/Phim/src/structs/apiServer"
	comite "github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/gin-gonic/gin"
)

func Initialization(comiteServer *comite.ComiteServer) *apiServer.ApiServer {

	router := gin.Default()

	return &apiServer.ApiServer{
		Route:         router,
		Comite_Server: (*comiteServer),
	}

}
