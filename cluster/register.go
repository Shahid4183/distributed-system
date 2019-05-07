package cluster

import (
	"math/rand"
	"net"
	"time"

	nodes "github.com/Shahid418/distributed-system/node"
)

// Connection - connection between two nodes
type Connection struct {
	ID          int
	Source      *nodes.Node
	Destination *nodes.Node
	Conn        net.Conn
}

// Register - registery for connections in our distributed system
type Register struct {
	Connections []Connection
}

// MakeRegister - creates a new registery
func MakeRegister() *Register {
	return &Register{}
}

// NewConnection - creates and returns new connection object
func NewConnection(s, d *nodes.Node, con net.Conn) Connection {
	rand.Seed(time.Now().UTC().UnixNano())
	return Connection{
		ID:          rand.Intn(100),
		Source:      s,
		Destination: d,
		Conn:        con,
	}
}

// Entry - make an entry of connection in registery
func (r *Register) Entry(s, d *nodes.Node, con net.Conn) {
	r.Connections = append(
		r.Connections,
		NewConnection(s, d, con),
	)
}
