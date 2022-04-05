package main

import (
	"github.com/Marinho3104/Phim/src/primatives"
	"github.com/Marinho3104/Phim/src/structs/comite"
)

func main() {
	comiteServer := make(chan *comite.ComiteServer)

	go primatives.ComiteServerOperational(comiteServer)

	primatives.ApiOperational(<-comiteServer)
}
