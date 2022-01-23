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
}

//each node will be a parent
//it will need 2 pointers, one for each left and right child
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	n.moveChildLeft(k)
	n.moveChildRight(k)
}

type MoveChild interface {
	moveChildLeft()
	moveChildRight()
}

func (n *Node) moveChildRight(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	}
}
func (n *Node) moveChildLeft(k int) {
	if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
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
