# Distributed System
**Aim:**

We have five nodes in a distributed system node-A, node-B, node-C, node-D, node-E. Design and implement a system in which one master node among the given nodes distribute the work to the slave nodes. After processing the data, the slave node returns data back to the master node which writes it to console.

**Explanation:**

At one moment of time, there can only be one master node. All nodes other than master node are called slave nodes.

**Example:**

Sample work can be sorting a list of names. For simplicity, you can read 1000 names from a file and sort the list by distributing the chunks of the list to slave nodes in the system. 

For simplicity, we can assume that all different nodes are on the same physical system running on different ports.
The master node may or may not participate in sorting the list(work).

**Evaluation:**

Write testable, maintainable and modular code
Plain TCP is preferred over HTTP for communication between nodes.

**Optional:**

Fault-tolerant design is plus i.e when one node goes down other can take its position.

## Code Structure

#### package cluster
Package **_cluster_** provides a way to make the cluster of nodes in our distributed computing environment. It has different types, constants and functions to achieve clusturing of nodes.


MAXNODES - defines max nodes our cluster can hold
```
const MAXNODES = 5
```
 PORT - defines start digit of port number
```
const PORT = "300"
```
Cluster - represents a cluster of nodes in our distributed system
```
type Cluster [MAXNODES]*nodes.Node
```
Create - creates a new cluster
This function will create MAXNODES number of nodes, starts them on port range starting from {PORT}1 to {PORT}{MAXNODES} adds them to the cluster and returns the newly created cluster
```
Create(wg *sync.WaitGroup) Cluster
```
> for example
> MAXNODE = 5
> PORT = 300
> cluster.Create(wg *sync.WaitGroup) will create 5 nodes on localhost using port 3001,3002,.....3005

GetMaster - randomely picks a master node from the cluster
Then stores reference of all other nodes in cluster as slave nodes in its Slave property
```
func (c Cluster) GetMaster() (*nodes.Node, error)
```


#### package node

Data - structure of data packet sent via tcp network between nodes
```
type Data []string
```

Node - holds tcp listener metadata about a node in distributed system
```
type Node struct {
	listener  net.Listener
	Slaves    []*Node
	ID        int    `json:"id"`
	IPAddress string `json:"ip_Address"`
	Port      string `json:"port"`
}
```
String - implement stringer interface and pretty printing the node info
```
(n Node) String() string
```
GetWaitGroup - creates and returns wait group to handle go routines of every node from main thread. Whenever a node is created, a new goroutine starts listening for its incomming connections. For this, we need a waitgroup to simulate communication between all go routines
```
GetWaitGroup() *sync.WaitGroup
```
MakeNode - creates a new node and start listening on given id, address and port
```
MakeNode(id int, port, ipAddress string, wg *sync.WaitGroup) (*Node, error)
```
Accept - waits for and returns the next connection to the listener
```
(n Node) Accept() (net.Conn, error)
```
Close - closes the listener
Any blocked Accept operations will be unblocked and return errors.
```
(n Node) Close() error
```
Addr - returns the listener's network address
```
func (n Node) Addr() net.Addr
```
listenOnPort - this method of type node will start listning on port of node n as soon as the node is created, in a different goroutine it will signal if it does its work via the work group reference variable to main process
```
(n Node) listenOnPort(wg *sync.WaitGroup)
```

#### package sort
QuickSort - implements sorting using quick sort algorithm
```
QuickSort(arr []string, low, high int)
```
Merge - merges left and right slice into newly created slice
```
Merge(left, right []string) []string
```
> merge function will merge two sorted arrays returning from slave nodes