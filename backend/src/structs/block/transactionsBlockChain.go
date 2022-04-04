package block

import "github.com/Marinho3104/Phim/src/structs/transaction"

type TransactionBlock struct {
	Id               int                       `json:"id"`
	TransactionCount int                       `json:"transactionCount"`
	Data             []transaction.Transaction `json:"data"`
	SignAddress      string                    `json:"signAddress"`
	Sign             string                    `json:"sign"`
}
