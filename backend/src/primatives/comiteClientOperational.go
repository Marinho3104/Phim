package primatives

import (
	work "github.com/Marinho3104/Phim-BlockChain/src/comiteClient/Work"
	maincomiteclientfunctions "github.com/Marinho3104/Phim-BlockChain/src/comiteClient/mainComiteClientFunctions"
)

func ComiteClientOperational() {
	comiteClient := maincomiteclientfunctions.Initialization("", "54321")

	maincomiteclientfunctions.SendContractOwnershipProve(comiteClient)

	work.GetWork(comiteClient)
}
