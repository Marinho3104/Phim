package maincomiteclient

import "github.com/Marinho3104/Phim/src/structs/comite"

func SendContractProof(comiteClient *comite.ComiteClient) {

	comiteClient.Connection.Write([]byte("testaddress"))

}
