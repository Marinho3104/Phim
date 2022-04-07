package contractfunctions

import (
	"errors"
	"os"
)

func DeleteContractsFile() error {

	path, err := os.Getwd()

	if err != nil {
		return errors.New("ENABLE TO GET CURRENT PATH")
	}

	os.Remove(path + "\\playground\\contractExecuter.py")
	os.Remove(path + "\\playground\\contract.py")

	return nil
}
