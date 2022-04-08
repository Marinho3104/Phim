package getblockchainfromfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Marinho3104/Phim/src/structs/block"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func GetBlockChainFromFile(comiteServer *comite.ComiteServer) (int, error) {

	_len := -1

	path, err := os.Getwd()
	if err != nil {
		return _len, errors.New("UNABLE TO GET CURRENT PATH")
	}

	files, _ := ioutil.ReadDir(path + "\\data\\transactionBlockChain")

	for i := 0; i < len(files); i++ {

		file, err := os.Open(path + fmt.Sprintf("\\data\\transactionBlockChain\\%d", i))

		if err != nil {
			return _len, fmt.Errorf("UNABLE TO OPEN %s FILE", path+fmt.Sprintf("\\data\\transactionBlockChain\\%d", i))
		}

		b, err := ioutil.ReadAll(file)

		if err != nil {
			return _len, fmt.Errorf("UNABLE TO READ THE CONTENT OF %d FILE", i)
		}

		_transactionBlock := block.TransactionBlock{}

		json.Unmarshal(b, &_transactionBlock)

		_len++

		comiteServer.BlockChainTransaction <- append(<-comiteServer.BlockChainTransaction, _transactionBlock)

	}

	files, _ = ioutil.ReadDir(path + "\\data\\contractBlockChain")

	for i := 0; i < len(files); i++ {

		file, err := os.Open(path + fmt.Sprintf("\\data\\contractBlockChain\\%d", i))

		if err != nil {
			return _len, fmt.Errorf("UNABLE TO OPEN %s FILE", path+fmt.Sprintf("\\data\\contractBlockChain\\%d", i))
		}

		b, err := ioutil.ReadAll(file)

		if err != nil {
			return _len, fmt.Errorf("UNABLE TO READ THE CONTENT OF %d FILE", i)
		}

		_contractBlock := block.ContractBlock{}

		json.Unmarshal(b, &_contractBlock)

		comiteServer.BlockChainContract <- append(<-comiteServer.BlockChainContract, _contractBlock)

	}

	files, _ = ioutil.ReadDir(path + "\\data\\contractInteractionsBlockChain")

	for i := 0; i < len(files); i++ {

		file, err := os.Open(path + fmt.Sprintf("\\data\\contractInteractionsBlockChain\\%d", i))

		if err != nil {
			return _len, fmt.Errorf("UNABLE TO OPEN %s FILE", path+fmt.Sprintf("\\data\\contractInteractionsBlockChain\\%d", i))
		}

		b, err := ioutil.ReadAll(file)

		if err != nil {
			return _len, fmt.Errorf("UNABLE TO READ THE CONTENT OF %d FILE", i)
		}

		_contractInteractionBlock := block.ContractInteractionsBlock{}

		json.Unmarshal(b, &_contractInteractionBlock)

		comiteServer.BlockChainContractInteractions <- append(<-comiteServer.BlockChainContractInteractions, _contractInteractionBlock)

	}

	return _len, nil
}
