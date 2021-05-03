// Breadth First Search algorithm
//
// Bit more complicated than dfs.go
// main difference is that instead of immediately traversing
// the next child we find, we want to add the values of every
// child for a single node before moving to next layer
// All this is done by utilizing a `queue`, that stores refs
// to nodes starting with a single struct. Then we start a for loop
// that continues as long as we have items in the queue. Since the queue
// represents the order of nodes left for us to traverse, we know that
// we'll have gone through it all when the queue is empty.
// The outer for loop represents the level of the graph we are on.

package main

type Node struct {
	Value    int
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []int) []int {
	queue := []*Node{n}

	for len(queue) > 0 {
		current := queue[0]
		queue := queue[1:]
		array = append(array, current.Value)
		for _, child := range n.Children {
			queue := append(queue, child)
		}
	}
}
