package contract

type Contract struct {
	CreatorAddress        string `json:"creatorAddress"`
	ContractAddress       string `json:"contractAddress"`
	ContractName          string `json:"contractName"`
	Disable               bool   `json:"disable"`
	AutoExec              bool   `json:"autoExec"`
	Info                  bool   `json:"info"`
	Data                  []byte `json:"data"`
	C                     int    `json:"c"`
	Sign                  string `json:"sign"`
	Fee                   int    `json:"fee"`
	ComiteReviewerAddress string `json:"comiteReviewerAddress"`
}
