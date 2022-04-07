package common

import (
	"fmt"
	"os"
)

func GetCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	return path

}
