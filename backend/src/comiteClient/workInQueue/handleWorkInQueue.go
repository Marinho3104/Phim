package workinqueue

import (
	"bytes"

	"github.com/Marinho3104/Phim/src/structs/comite"
)

func HandleWorkInQueue(comiteClient *comite.ComiteClient) {

	for {
		_workInQueue := <-comiteClient.WokInQueue

		respSplit := bytes.Split(_workInQueue)

	}

}
