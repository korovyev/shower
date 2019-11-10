package main

import "fmt"

// TorrentClient will do most stuff presumably - TODO document.
type TorrentClient struct {
	// maps a peer to its status (false for no connection / true for connected peer, will need expanding)
	peers map[Peer]bool
}

func (t *TorrentClient) handleNewPeer(newPeer chan Peer) {
	// limiting peers for now
	for len(t.peers) < 100 {
		peer := <-newPeer

		fmt.Println(peer.ip, peer.port)

		_, exists := t.peers[peer]

		if !exists {
			t.peers[peer] = false
			fmt.Println("wahoo!", len(t.peers))
		} else {
			fmt.Println("not a new peer")
		}
	}
}
