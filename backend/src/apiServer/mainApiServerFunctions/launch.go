package mainapiserverfunctions

import "github.com/Marinho3104/Phim/src/structs/apiServer"

func Launch(_apiServer *apiServer.ApiServer) {

	_apiServer.Route.Run("localhost:8000")

}
