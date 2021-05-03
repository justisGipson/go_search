// depth-first search algorithm
//
// NOTE: NOT A RECURSIVE ALGORITHM!
// Just calling DFS on each child node we unpack
// `range` is used to loop over a node's children
// then discard the 1st arg with _, it's an index we
// don't care about

package main

type Node struct {
	value    int
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []int) []int {
	array = append(array, n.value)
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}
}
