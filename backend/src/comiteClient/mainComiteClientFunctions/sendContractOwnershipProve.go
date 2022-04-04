package maincomiteclientfunctions

import "github.com/Marinho3104/Phim-BlockChain/src/structs/comite"

func SendContractOwnershipProve(comiteClient *comite.ComiteClient) {

	comiteClient.Connection.Write([]byte("testaddress"))

}
