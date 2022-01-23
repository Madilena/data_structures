package main

import "fmt"

var count int

func main() {
	//tree has to be an address
	tree := &Node{Key: 100}
	array := []int{100, 3, 25, 29, 52, 1, 8, 9, 11, 18, 19, 69, 75, 92, 22}
	for i := 0; i < len(array); i++ {
		tree.Insert(array[i])
	}
	fmt.Println(tree)
	fmt.Println(tree.Search(75))
	fmt.Println(count)
	fmt.Println(maxDepth(tree))
}

//each node will be a parent
//it will need 2 pointers, one for each left and right child
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.keyIsGreaterThanParent(k) {
		n.moveKeyToRightChild(k)
	} else if n.keyIsSmallerThanParent(k) {
		n.moveKeyToLeftChild(k)
	}
}

type MoveChild interface {
	moveKeyToLeftChild()
	moveKeyToRightChild()
	keyIsSmallerThanParent()
	keyIsGreaterThanParent()
}

func maxDepth(n *Node) int {
	if n == nil {
		return -1
	}
	lDepth := maxDepth(n.Left)
	rDepth := maxDepth(n.Right)

	if lDepth > rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}

func (n *Node) keyIsSmallerThanParent(k int) bool {
	//n.Key == parent
	return k < n.Key
}

func (n *Node) keyIsGreaterThanParent(k int) bool {
	//n.Key == parent
	return k > n.Key
}

func (n *Node) moveKeyToRightChild(k int) {
	if n.Right == nil {
		n.Right = &Node{Key: k}
	} else {
		n.Right.Insert(k)
	}
}
func (n *Node) moveKeyToLeftChild(k int) {
	if n.Left == nil {
		n.Left = &Node{Key: k}
	} else {
		n.Left.Insert(k)
	}
}

func (n *Node) Search(k int) bool {
	count++
	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}
