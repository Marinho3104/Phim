package contract

import "github.com/Marinho3104/Phim/src/structs/transaction"

type ContractInteraction struct {
	ContractAddress       string                  `json:"contractAddress"`
	Transaction           transaction.Transaction `json:"transaction"`
	Function              string                  `json:"function"`
	Arguments             map[string]interface{}  `json:"arguments"`
	Variables             map[string]interface{}  `json:"variables"`
	Fee                   int                     `json:"fee"`
	Sign                  string                  `json:"sign"`
	C                     int                     `json:"c"`
	ComiteReviewerAddress string                  `json:"comiteReviewerAddress"`
}
