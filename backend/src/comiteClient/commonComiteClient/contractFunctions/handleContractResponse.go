package contractfunctions

import (
	"encoding/json"
	"errors"
)

func HandleContractResponse(resp []byte) (map[string]interface{}, error) {

	_response := make(map[string]interface{})

	err := json.Unmarshal(resp, &_response)

	if err != nil {

		return make(map[string]interface{}), errors.New("RESPONSE STRUCT FROM CONTRACT EXECUTION WRONG")

	}
	return _response, nil

}
