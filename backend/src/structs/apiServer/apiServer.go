package apiServer

import (
	comite "github.com/Marinho3104/Phim/src/structs/comite"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	Route         *gin.Engine
	Comite_Server *comite.ComiteServer
}
