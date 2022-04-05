package transaction

import (
	comiteclientconnection "github.com/Marinho3104/Phim/src/structs/comite/comiteClientConnection"
)

type ConfirmationTransaction struct {
	Hash256              string
	Transation           Transaction
	ComiteChoose         comiteclientconnection.ComiteClientConnection
	ComiteChooseResponse int
	ComiteTotal          []comiteclientconnection.ComiteClientConnection
	ComiteAccepted       []comiteclientconnection.ComiteClientConnection
	ComiteDeclined       []comiteclientconnection.ComiteClientConnection
}
