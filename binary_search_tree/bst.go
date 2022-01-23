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
	fmt.Println("75 exists in tree:", tree.Search(75))
	fmt.Println("We checked this many nodes to find it:", count)
	fmt.Println("This is the height of the tree:", maxDepth(tree))

	fmt.Println("inOrder traversal")
	inOrder(tree)
	fmt.Println("preOrder traversal")
	preOrder(tree)
	fmt.Println("postOrder travesal")
	postOrder(tree)
	fmt.Println("levelOrder")
	levelOrder(tree)
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

//Inorder traversal gives nodes in non-decreasing order.
func inOrder(n *Node) {
	if n == nil {
		return
	}
	inOrder(n.Left)
	fmt.Println(n.Key)
	inOrder(n.Right)
}

//Preorder traversal is used to create a copy of the tree.
//Preorder traversal is also used to get prefix expression on an expression tree.
//    eg The expression for adding the numbers 1 and 2 is written in Polish notation as + 1 2 (pre-fix), rather than as 1 + 2 (in-fix).
func preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.Key)
	inOrder(n.Left)
	inOrder(n.Right)
}

//Postorder traversal is used to delete the tree.
//Postorder traversal is also useful to get the postfix expression of an expression tree.
func postOrder(n *Node) {
	if n == nil {
		return
	}
	inOrder(n.Left)
	inOrder(n.Right)
	fmt.Println(n.Key)
}

func levelOrder(n *Node) {
	for level := 0; level < maxDepth(n); level++ {
		printCurrentLevel(n, level)
	}
}

func printCurrentLevel(n *Node, level int) {
	if n == nil {
		return
	}
	if level == 1 {
		fmt.Print(n.Key)
	} else if level > 1 {
		printCurrentLevel(n.Left, level-1)
		printCurrentLevel(n.Right, level-1)
	}
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
