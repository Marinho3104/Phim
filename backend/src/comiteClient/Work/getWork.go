package work

import (
	"fmt"
	"os"
	"strings"

	"github.com/Marinho3104/Phim-BlockChain/src/structs/comite"
)

func GetWork(comiteClient *comite.ComiteClient) {

	all := []byte("")

	fmt.Println("Getting work")

	for {
		resp := make([]byte, 150000)

		_len, err := comiteClient.Connection.Read(resp)

		if err != nil {
			fmt.Println("Server response error ... \nAborting")
			os.Exit(1)
		}

		resp = resp[:_len]

		all = append(all, resp...)

		respSplit := strings.Split(string(all), "\n")

		all = HandleWork(comiteClient, respSplit)
	}

}
