package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sync"

	"github.com/Shahid418/distributed-system/cluster"
	nodes "github.com/Shahid418/distributed-system/node"
	"github.com/Shahid418/distributed-system/sort"
)

func main() {
	// create a wait group for goroutines
	wg := nodes.GetWaitGroup()
	// create a cluster
	mycluster := cluster.Create(wg)
	// get master from the cluster
	// GetMaster method of cluster picks and returns a random
	// node from the cluster
	master, err := mycluster.GetMaster()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// make registery of connections of nodes
	register := cluster.MakeRegister()
	// make connections of master node to all slave nodes
	// store the connections in the registery
	for _, slave := range master.Slaves {
		// dial on slave node's ip and port
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", slave.IPAddress, slave.Port))
		if err != nil {
			fmt.Println("Cannot establish connection to node ", slave.ID)
		}
		// make an entry in registry
		register.Entry(master, slave, conn)
	}
	fmt.Println("Simulating sorting task in our distributed system")
	fmt.Println("------------------------------------------")
	// simulate sorting task
	SimulateSortingTask(master, register)
	fmt.Println("------------------------------------------")
	fmt.Println("Sort task completed successfully")
	fmt.Println("------------------------------------------")
	// wait for go routines to finish
	wg.Wait()
}

// SimulateSortingTask - this function is used to simulate the distributed sorting task
func SimulateSortingTask(m *nodes.Node, reg *cluster.Register) {
	var wg sync.WaitGroup
	// open file
	f, err := os.Open("city_unsorted.json")
	if err != nil {
		fmt.Println("Unable to open file")
	}
	// read from file
	byteData, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Error while reading file")
	}
	// variable to store unmarshaled data
	var data nodes.Data
	// unmarshal data into object
	if err := json.Unmarshal(byteData, &data); err != nil {
		fmt.Println("Error while unmarshalling byte data")
	}
	// send chunks of data to slave nodes
	chunkSize := len(data) / len(m.Slaves)
	var sortedList []string
	for i, con := range reg.Connections {
		c := make(chan []string)
		if err := json.NewEncoder(con.Conn).Encode(data[i*chunkSize : chunkSize*(i+1)]); err != nil {
			fmt.Println("Error while sending data to slave node\n", err)
			continue
		}
		// wait for response
		wg.Add(1)
		go waitForResponse(con.Destination.ID, con.Conn, &wg, c)
		// get response and merge it with sorted list
		sortedList = sort.Merge(sortedList, <-c)
	}
	for _, city := range sortedList {
		fmt.Println(city)
	}
	// wait for all go routines to complete
	wg.Wait()
}

func waitForResponse(id int, c net.Conn, wg *sync.WaitGroup, ch chan []string) {
	var resp nodes.Data
	if err := json.NewDecoder(c).Decode(&resp); err != nil {
		fmt.Println("Error while receiving data from slave node\n", err)
	}
	// send response from slave nodes to master
	ch <- resp
	// signal wait group that this go routine has finished its job
	wg.Done()
}
