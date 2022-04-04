package transaction

type Transaction struct {
	AddressFrom           string `json:"addressFrom"`
	AddressTo             string `json:"addressTo"`
	Amount                int    `json:"amount"`
	C                     int    `json:"c"`
	Sign                  string `json:"sign"`
	Fee                   int    `json:"fee"`
	ComiteReviewerAddress string `json:"comiteReviewerAddress"`
}
