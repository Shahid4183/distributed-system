package cluster

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	nodes "github.com/Shahid418/distributed-system/node"
)

// MAXNODES - defines max nodes our cluster can hold
const MAXNODES = 5

// PORT - defines start digit of port number
const PORT = "300"

// Cluster - represents a cluster of nodes in our distributed system
type Cluster [MAXNODES]*nodes.Node

// Create - creates a new cluster
func Create(wg *sync.WaitGroup) Cluster {
	fmt.Println("Creating cluster")
	fmt.Println("------------------------------------------")
	// create a cluster object
	var cluster Cluster
	// create N nodes where
	// N = MAXNODES
	// and add the nodes to the cluster
	for i := 0; i < MAXNODES; i++ {
		// call node package's MakeNode method
		// pass port and ip address
		n, err := nodes.MakeNode(
			i+1,
			fmt.Sprintf("%s%d", PORT, i),
			"localhost",
			wg,
		)
		if err != nil {
			fmt.Printf("Error while creating cluster:%+v\n", err)
		}
		cluster[i] = n
	}
	fmt.Println("Cluster created successfully")
	fmt.Println("------------------------------------------")
	return cluster
}

// GetMaster - randomely picks a master node from the cluster
func (c Cluster) GetMaster() (*nodes.Node, error) {
	// using package rand to generate random number
	// seeding current unix timestamp to rand package
	rand.Seed(time.Now().UTC().UnixNano())
	// selecting a random node as master node from clsuter
	masterNodeNo := rand.Intn(MAXNODES)
	master := c[masterNodeNo]
	// other nodes will become salve nodes and establishes connection with master node
	for i, n := range c {
		if i == masterNodeNo {
			continue
		}
		master.Slaves = append(master.Slaves, n)
	}
	// return the master nodes
	return master, nil
}
