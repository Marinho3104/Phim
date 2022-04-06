package primatives

import (
	maincomiteclientfunctions "github.com/Marinho3104/Phim/src/comiteClient/mainComiteClientFunctions"
	work "github.com/Marinho3104/Phim/src/comiteClient/work"
)

func ComiteClientOperational() {
	comiteClient := maincomiteclientfunctions.Initialization("", "54321")

	maincomiteclientfunctions.SendContractOwnershipProve(comiteClient)

	go work.HandleWork(comiteClient)

	work.GetWork(comiteClient)
}
