package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// TorrentClient will do most stuff presumably - TODO document.
type TorrentClient struct {
	// maps a peer to its status (false for no connection / true for connected peer, will need expanding)
	peers map[Peer]PeerConnectionStatus

	infoHash string
}

func (t *TorrentClient) handleNewPeer(newPeer chan Peer) {
	// limiting peers for now
	for len(t.peers) < 1 {
		peer := <-newPeer

		_, exists := t.peers[peer]

		if !exists {
			t.peers[peer] = New
		}
	}
}

func (t *TorrentClient) connectToPeers() {
	for {
		new := make([]Peer, 0)
		connecting := make([]Peer, 0)
		for peer, connectionStatus := range t.peers {
			switch connectionStatus {
			case New:
				new = append(new, peer)
				t.peers[peer] = Connecting
				t.connect(peer)
			case Connecting:
				connecting = append(connecting, peer)
			default:
				fmt.Println("nope")
			}
		}

		// fmt.Println("NEW: ", len(new), " -- CONNECTING: ", len(connecting))
	}
}

func (t *TorrentClient) connect(peer Peer) {

	fmt.Println("\n=====HANDSHAKE=====")

	add := []string{
		peer.ip,
		peer.port,
	}

	fullIP := strings.Join(add, ":")

	// handshake format:
	// chr(19)+"BitTorrent Protocol"+8*chr(0)+info_hash+peer_id
	// handshake := "19BitTorrent Protocol00000000" + t.infoHash + "-SH0000-123456789012"

	id := "-SH0000-123456789012"
	kBitTorrentHeader := []byte{'\x13', 'B', 'i', 't', 'T', 'o', 'r', 'r', 'e', 'n', 't', ' ', 'p', 'r', 'o', 't', 'o', 'c', 'o', 'l'}
	handshake := make([]byte, 68)

	// copy(handshake[:1], "19")
	copy(handshake, kBitTorrentHeader[0:])
	copy(handshake[28:48], []byte(t.infoHash))
	copy(handshake[48:68], []byte(id))

	fmt.Println("Attempting handshake to: ", fullIP, "with: ", string(handshake))

	connection, error := net.Dial("tcp", fullIP)

	if error != nil {
		log.Fatal(error)
	}

	reader := bufio.NewReader(connection)
	writer := bufio.NewWriter(connection)

	size, errr := writer.Write(handshake)
	writer.Flush()

	if errr != nil {
		log.Fatal(errr)
	}
	fmt.Println("bytes sent: ", size)

	for {
		str, err := reader.ReadString('\n')
		if len(str) > 0 {
			fmt.Println(str)
		}
		if err != nil {
			fmt.Println("fatal!")
			log.Fatal(err)
		}
	}

}
