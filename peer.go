package main

// PeerConnectionStatus models the connection state between ourselves and a giuven Peer
type PeerConnectionStatus int32

// I miss enums
const (
	New        PeerConnectionStatus = iota
	Connecting PeerConnectionStatus = iota
	Dropped    PeerConnectionStatus = iota
	Connected  PeerConnectionStatus = iota
)

// Peer represents another client in the swarm
type Peer struct {
	ip   string
	port string
}
