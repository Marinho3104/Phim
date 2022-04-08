package block

import "github.com/Marinho3104/Phim/src/structs/contract"

type ContractInteractionsBlock struct {
	Id                       int                            `json:"id"`
	ContractInteractionCount int                            `json:"contractIntractionCount"`
	Data                     []contract.ContractInteraction `json:"data"`
	SignAddress              string                         `json:"signAddress"`
	Sign                     string                         `json:"sign"`
}
