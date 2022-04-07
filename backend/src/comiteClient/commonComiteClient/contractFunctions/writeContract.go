package contractfunctions

import (
	"fmt"
	"os"
	"strings"

	"github.com/Marinho3104/Phim/src/structs/contract"
)

func WriteContract(_contract contract.Contract, contractBalance int, interactor string, interactorAmount int, fee int, variables map[string]interface{}, creation bool) {

	code := string(_contract.Data)

	code += initDef(_contract.ContractAddress, interactor, contractBalance, fee, interactorAmount, variables, creation)

	f, err := os.OpenFile("playground/contract.py", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(code))
	if err != nil {
		panic(err)
	}

	f.Close()

}

func initDef(contractAddress, interactor string, contractBalance, fee, interactorAmount int, variables map[string]interface{}, creation bool) (_code string) {

	_code += "\n    def __init__(self) -> None:"

	_code += fmt.Sprintf("\n        super().__init__(\"%s\", %d, %d, %s)", fixInvalidCharacters(contractAddress), contractBalance, fee, getInteractorInfo(interactor, interactorAmount))

	_code += setVariables(variables)

	if creation {
		_code += "\n        self.initialization()"
	}

	return
}

func setVariables(variables map[string]interface{}) (variablesString string) {

	for s, v := range variables {
		switch value := v.(type) {
		case int:
			variablesString += fmt.Sprintf("\n        self.%s = %d", s, v)
		case string:
			v = strings.Replace(value, "\\", "\\\\", -1)

			v = strings.Replace(value, "\"", "\\\"", -1)

			variablesString += fmt.Sprintf("\n        self.%s = \"%s\"", s, value)
		case float64:
			variablesString += fmt.Sprintf("\n        self.%s = %g", s, v)
		case []interface{}:
			variablesString += fmt.Sprintf("\n        self.%s = [ %s ]", s, SetArray(value))
		default:
			variablesString += fmt.Sprintf("\n        self.%s = %s", s, v)
		}
	}

	return
}

func SetArray(arr []interface{}) (variablesString string) {

	for _, v := range arr {
		switch value := v.(type) {
		case int:
			variablesString += fmt.Sprintf("%d, ", v)
		case string:
			v = strings.Replace(value, "\\", "\\\\", -1)

			v = strings.Replace(value, "\"", "\\\"", -1)

			variablesString += fmt.Sprintf("\"%s\", ", value)
		case float64:
			variablesString += fmt.Sprintf("%g, ", v)
		case []interface{}:
			variablesString += fmt.Sprintf("[ %s ]", SetArray(value))
		default:
			variablesString += fmt.Sprintf("%s, ", v)
		}
	}

	variablesString = variablesString[:len(arr)-2]

	return
}

func getInteractorInfo(interactorAddress string, amount int) (interactorInfoCode string) {

	interactorInfoCode += fmt.Sprintf("PhimChainContract.InteractorInfo(\"%s\", %d)", fixInvalidCharacters(interactorAddress), amount)

	return
}

func fixInvalidCharacters(_string string) string {

	_string = strings.Replace(_string, "\\", "\\\\", -1)

	_string = strings.Replace(_string, "\"", "\\\"", -1)

	return _string
}
