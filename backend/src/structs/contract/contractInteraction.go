package contract

type ContractInteraction struct {
	ContractAddress       string                 `json:"contractAddress"`
	Interactor            string                 `json:"interactor"`
	Amount                int                    `json:"amount"`
	Function              string                 `json:"function"`
	Arguments             map[string]interface{} `json:"arguments"`
	Variables             map[string]interface{} `json:"variables"`
	Fee                   int                    `json:"fee"`
	Sign                  string                 `json:"sign"`
	C                     int                    `json:"c"`
	ComiteReviewerAddress string                 `json:"comiteReviewerAddress"`
}
