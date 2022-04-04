package getblockchainfromfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Marinho3104/Phim-BlockChain/src/structs/block"
	"github.com/Marinho3104/Phim-BlockChain/src/structs/comite"
)

func GetBlockChainFromFile(comiteServer *comite.ComiteServer) error {

	path, err := os.Getwd()
	if err != nil {
		return errors.New("UNABLE TO GET CURRENT PATH")
	}

	files, _ := ioutil.ReadDir(path + "\\data\\transactionBlockChain")

	for i := 0; i < len(files); i++ {

		file, err := os.Open(path + fmt.Sprintf("\\data\\transactionBlockChain\\%d", i))

		if err != nil {
			return fmt.Errorf("UNABLE TO OPEN %s FILE", path+fmt.Sprintf("\\data\\transactionBlockChain\\%d", i))
		}

		b, err := ioutil.ReadAll(file)

		if err != nil {
			return fmt.Errorf("UNABLE TO READ THE CONTENT OF %d FILE", i)
		}

		_transactionBlock := block.TransactionBlock{}

		json.Unmarshal(b, &_transactionBlock)

		comiteServer.BlockChainTransaction <- append(<-comiteServer.BlockChainTransaction, _transactionBlock)

	}

	return nil
}
