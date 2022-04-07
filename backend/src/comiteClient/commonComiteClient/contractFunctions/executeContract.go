package contractfunctions

import (
	"errors"
	"os"
	"os/exec"
	"time"
)

func ExecuteContract() ([]byte, error) {

	path, err := os.Getwd()

	if err != nil {
		return nil, errors.New("ENABLE TO GET CURRENT PATH")
	}

	resp, _ := exec.Command(path+"\\playground\\python.exe", path+"\\playground\\contractExecuter.py").Output()

	initialTime := time.Now()

	for {
		if time.Since(initialTime).Seconds() > 5.0 || string(resp) != "" {
			break
		}
	}

	err = nil

	if string(resp) == "" {
		err = errors.New("NO RESPONSE FROM CONTRACT OR ERROR EXECUTING")
	}

	return resp, err

}
