package nodes

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"github.com/Shahid418/distributed-system/sort"
)

// Data - structure of data packet sent via tcp network between nodes
type Data []string

// Node - holds tcp listener metadata about a node in distributed system
type Node struct {
	listener  net.Listener
	Slaves    []*Node
	ID        int    `json:"id"`
	IPAddress string `json:"ip_Address"`
	Port      string `json:"port"`
}

// String - implement stringer interface and pretty printing the node info
func (n Node) String() string {
	// using json's marshal indent to format and return node information
	s, _ := json.MarshalIndent(n, "", "\t")
	return string(s)
}

// GetWaitGroup - gets wait group
func GetWaitGroup() *sync.WaitGroup {
	return new(sync.WaitGroup)
}

// MakeNode - creates a new node and start listening on given id address and port
func MakeNode(id int, port, ipAddress string, wg *sync.WaitGroup) (*Node, error) {
	fmt.Println("Creating Node:", id)
	fmt.Println("------------------------------------------")
	// create a tcp listener
	// using given ip address and port
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ipAddress, port))
	// return error if any
	if err != nil {
		return nil, err
	}
	// initialize and return new node
	// nil for no error
	n := &Node{
		ID:        id,
		Port:      port,
		IPAddress: ipAddress,
		listener:  l,
	}
	wg.Add(1)
	go n.listenOnPort(wg)
	fmt.Printf("Node %d created successfully and started listening on port %s\n", id, port)
	fmt.Println("Node Info:", n)
	fmt.Println("------------------------------------------")
	return n, nil
}

// Accept - waits for and returns the next connection to the listener
func (n Node) Accept() (net.Conn, error) {
	return n.listener.Accept()
}

// Close - closes the listener
// Any blocked Accept operations will be unblocked and return errors.
func (n Node) Close() error {
	return n.listener.Close()
}

// Addr - returns the listener's network address.
func (n Node) Addr() net.Addr {
	return n.listener.Addr()
}

func (n Node) listenOnPort(wg *sync.WaitGroup) {
	// accept incoming connection
	con, err := n.Accept()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// create data object to receive data from master node
		var data Data
		json.NewDecoder(con).Decode(&data)
		// sort data using quick sort
		sort.QuickSort(data, 0, len(data)-1)
		// send sorted array back to master
		json.NewEncoder(con).Encode(data)
	}
	wg.Done()
}
