package primatives

import (
	mainapiserverfunctions "github.com/Marinho3104/Phim/src/apiServer/mainApiServerFunctions"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func ApiOperational(_comiteServer *comite.ComiteServer) {

	apiServer := mainapiserverfunctions.Initialization(_comiteServer)

	mainapiserverfunctions.AddRoutes(apiServer)

	mainapiserverfunctions.Launch(apiServer)

}
