package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nictuku/dht"
)

// PeerProducer produces peers - TODO: proper documenting
type PeerProducer struct {
	foundPeers chan Peer
}

const (
	httpPortTCP = 8711
)

func (p *PeerProducer) startDHT(infoHash string) {
	ih, err := dht.DecodeInfoHash(infoHash)
	if err != nil {
		fmt.Printf("DecodeInfoHash error: %v\n", err)
		os.Exit(1)
	}

	d, err := dht.New(nil)
	if err != nil {
		fmt.Printf("NewDHTNode error: %v", err)
		os.Exit(1)

	}
	// For debugging.
	go http.ListenAndServe(fmt.Sprintf(":%d", httpPortTCP), nil)

	if err = d.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "DHT start error: %v", err)
		os.Exit(1)
	}

	go p.drainresults(d)

	for {
		// Give the DHT some time to "warm-up" its routing table.
		time.Sleep(5 * time.Second)
		d.PeersRequest(string(ih), false)
	}
}

// drainresults loops, printing the address of nodes it has found.
func (p *PeerProducer) drainresults(d *dht.DHT) {
	fmt.Println("=========================== DHT")
	fmt.Printf("Note that there are many bad nodes that reply to anything you ask.")
	fmt.Printf("Peers found:")
	for r := range d.PeersRequestResults {
		for _, peers := range r {
			for _, x := range peers {
				addr := dht.DecodePeerAddress(x)
				addrWithPort := strings.Split(addr, ":")
				peer := Peer{addrWithPort[0], addrWithPort[1]}
				p.foundPeers <- peer
			}
		}
	}
}
