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
