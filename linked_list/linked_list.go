package main

import "fmt"

func main() {
	linkedList := NewLinkedList()
	fmt.Println("\n", "Remember the head is in the stack and the tail is in the heap", "\n")

	fmt.Println("\n", "Adding nodes to linked list", "\n")
	linkedList.Append(1).Append(2).Append(3).Append(4).Append(5).Append(6).PrintMemoryAddresses()

	fmt.Println("\n", "Just printing again", "\n")
	linkedList.PrintMemoryAddresses()

	fmt.Println("\n", "Now deleting node with value 2", "\n")
	linkedList.DeleteWithValue(2).PrintMemoryAddresses()

	fmt.Println("\n", "Now adding a node with value 7", "\n")
	linkedList.Append(7).PrintMemoryAddresses()
}

// LinkedList : Data structure
type LinkedList struct {
	Head *Node
	Tail *Node
}

// Node : A Linked List node with data of any type
// a node needs a value and a reference to another node
type Node struct {
	Data interface{}
	Next *Node
}

func (ll *LinkedList) PrintMemoryAddresses() {
	var node = ll.Head
	fmt.Println("head", node.Data, "has memory address", &node)
	for {
		if node.Next == nil {
			return
		}
		node = node.Next
		fmt.Println("node", node.Data, "has memory address", &node.Next)
	}
}

//constructor returns pointer to new instance of linked list
func NewLinkedList() *LinkedList {
	emptyNode := &Node{
		Next: nil,
		Data: nil,
	}
	return &LinkedList{
		Head: emptyNode,
		Tail: emptyNode,
	}
}

// Append : Appending a new node to the end of the Linked List
func (ll *LinkedList) Append(d interface{}) *LinkedList {
	nextNode := &Node{
		Next: nil,
		Data: d,
	}
	updateCurrentTail(ll, nextNode)
	//set the linked lists' tail to nextNode
	ll.Tail = nextNode
	return ll
}

//first need to handle pre-exinting node's head & tail
func updateCurrentTail(ll *LinkedList, nextNode *Node) {
	//if head is empty, set head to nextNode
	//otherwise, we set the current tail's Next value to our nextNode
	if ll.Head.Data == nil {
		ll.Head = nextNode
	} else {
		ll.Tail.Next = nextNode
	}
}

// DeleteWithValue : Deletes node which has a specific value
func (ll *LinkedList) DeleteWithValue(v interface{}) *LinkedList {
	var node = ll.Head
	//if the desired node to delete is at the head, set the head to its Next property
	if node.Data == v {
		ll.Head = ll.Head.Next
		return ll
	}

	//infinite loop
	for {
		//if the desired node to delete is the next node, set that node's Next property
		if isNextNodeDesiredNode(v, node) {
			if isNextNodeNotLast(node) {
				shiftNodesNextUp(node)
				break
			}
			//if the node to delete is second to last, the new node is now the end of linked list
			setNodeAsLast(node)
			break
		}
		shiftNextNodeUp(node)
	}
	return ll
}

//we're looking at current node's next property to see if we want to delete the next node
func isNextNodeDesiredNode(v interface{}, node *Node) bool {
	return v == node.Next.Data
}

//we're looking at the next node's next property to see if it is the last node in linked list
func isNextNodeNotLast(node *Node) bool {
	return node.Next.Next != nil
}

//our 'delete' operation is really just making the current node's pointer skip the node we want to delete
//and that pointer will point to the node after the node we want to delete
func shiftNodesNextUp(node *Node) {
	node.Next = node.Next.Next
}

//we're going to set the current node to the next node
func shiftNextNodeUp(node *Node) {
	node = node.Next
}

//the last node in the linked list need's to have nil as its next property
func setNodeAsLast(node *Node) {
	node.Next = nil
}
