package main

import (
	"flag"
)

func main() {
	magnetPointer := flag.String("magnet", "default", "Magnet link to use")
	flag.Parse()
	params := parse(*magnetPointer)

	xt := params["xt"]

	infoHash := infoHash(xt)
	startDHT(infoHash)
}
