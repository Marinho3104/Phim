package contract

import comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"

type ConfirmationContract struct {
	Hash256              string
	Contract             Contract
	ComiteChoose         comiteclientconnection.ComiteClientConnection
	ComiteChooseResponse int
	ComiteTotal          []comiteclientconnection.ComiteClientConnection
	ComiteResponse       []comiteclientconnection.ComiteClientConnection
	Resp                 map[string]int
}
