package handleworkfunctions

import (
	"bytes"
	"net"
)

func SendResponseToComiteServer(comiteServerConnection net.Conn, info []string) {

	comiteServerConnection.Write(
		bytes.Join(
			[][]byte{
				[]byte(info[0]),
				[]byte(info[1]),
				[]byte(""),
			},
			[]byte("\n"),
		),
	)

}
