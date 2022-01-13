package main

import "fmt"

func main() {
	linkedList := New()
	linkedList.Append(1).Append(2).Append(3).PrintAll()
}

// LinkedList : Data structure
type LinkedList struct {
	Head *Node
	Tail *Node
}

// Node : A Linked List node with data of any type
type Node struct {
	Next *Node
	Data interface{}
}

func (ll *LinkedList) PrintAll() {
	var node = ll.Head
	for {
		fmt.Println(node.Data)
		if node.Next == nil {
			return
		}
		node = node.Next
	}
}

//constructor returns pointer to new instance of linked list
func New() *LinkedList {
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
	updateCurrentNodes(ll, nextNode)
	//set the linked lists' tail to nextNode
	ll.Tail = nextNode
	return ll
}

//first need to handle pre-exinting node's head&tail then need to handle tail
func updateCurrentNodes(ll *LinkedList, nextNode *Node) {
	//if head is empty, set head to nextNode
	//otherwise, we set the current tail's Next value to our nextNode
	if ll.Head.Data == nil {
		ll.Head = nextNode
	} else {
		ll.Tail.Next = nextNode
	}
}