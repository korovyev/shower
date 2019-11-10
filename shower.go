package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	magnetPointer := flag.String("magnet", "default", "Magnet link to use")
	flag.Parse()

	params := parse(*magnetPointer)

	xt, ok := params["xt"]

	if !ok {
		log.Fatal("not a real link!", params)
	}

	infoHash := infoHash(xt)
	peerProducer := &PeerProducer{
		foundPeers: make(chan Peer),
	}
	go peerProducer.startDHT(infoHash)
	client := TorrentClient{
		peers: make(map[Peer]bool),
	}
	client.handleNewPeer(peerProducer.foundPeers)
}

func doStuff(chanchan chan int) {
	fmt.Println("waitin...")
	for {
		num := <-chanchan
		fmt.Println("success!", num)
	}
}
