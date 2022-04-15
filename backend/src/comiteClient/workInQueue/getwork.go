package workinqueue

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Marinho3104/Phim/src/structs/comite"
)

func GetWork(comiteClient *comite.ComiteClient) {

	fmt.Println("Getting work")

	for {

		resp := make([]byte, 150000)

		_len, err := comiteClient.Connection.Read(resp)

		if err != nil {
			fmt.Println("Server response error ... \nAborting")
			os.Exit(1)
		}

		resp = resp[:_len]

		select {

		case _currentWorkToDo := <-comiteClient.WorkInQueue:

			comiteClient.WorkInQueue <- bytes.Join(
				[][]byte{
					_currentWorkToDo,
					resp,
				},
				[]byte("\n"),
			)

		default:

			comiteClient.WorkInQueue <- resp

		}

	}

}
