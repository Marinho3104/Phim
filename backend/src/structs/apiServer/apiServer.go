package apiServer

import (
	comite "github.com/Marinho3104/Phim-BlockChain/src/structs/Comite"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	Route         *gin.Engine
	Comite_Server *comite.ComiteServer
}
