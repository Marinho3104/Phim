package block

import "github.com/Marinho3104/Phim/src/structs/contract"

type ContractBlock struct {
	Id            int                 `json:"id"`
	ContractCount int                 `json:"contractCount"`
	Data          []contract.Contract `json:"data"`
	SignAddress   string              `json:"signAddress"`
	Sign          string              `json:"sign"`
}
