package contractfunctions

import (
	"fmt"
	"os"
	"strings"

	"github.com/Marinho3104/Phim/src/structs/contract"
)

func WriteExecuter(_contract contract.Contract, function string, arguments map[string]interface{}) {

	_code := fmt.Sprintf("from contract import %s\n", _contract.ContractName)

	_code += fmt.Sprintf("%s()", _contract.ContractName)

	if function != "" {

		_code += fmt.Sprintf(".%s(", function)

		_code += setArguments(arguments)

		_code += ")"

	}

	f, err := os.OpenFile("playground/contractExecuter.py", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(_code))
	if err != nil {
		panic(err)
	}

	f.Close()

}

func setArguments(variables map[string]interface{}) (variablesString string) {

	if len(variables) == 0 {
		return
	}

	for s, v := range variables {
		switch value := v.(type) {
		case int:
			variablesString += fmt.Sprintf("%s = %d, ", s, v)
		case string:
			v = strings.Replace(value, "\\", "\\\\", -1)

			v = strings.Replace(value, "\"", "\\\"", -1)

			variablesString += fmt.Sprintf("%s = \"%s\", ", s, value)
		case float64:
			variablesString += fmt.Sprintf("%s = %g, ", s, v)
		case []interface{}:
			variablesString += fmt.Sprintf("%s = [ %s ] ,", s, SetArray(value))
		default:
			variablesString += fmt.Sprintf("%s = %s, ", s, v)
		}
	}

	variablesString = variablesString[:len(variablesString)-2]

	return
}
