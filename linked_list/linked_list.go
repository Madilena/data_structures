package main

import "fmt"

func main() {
	linkedList := NewLinkedList()
	linkedList.Append(1).Append(2).Append(3).PrintAll()
	linkedList.DeleteWithValue(2).PrintAll()
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
	updateCurrentNodes(ll, nextNode)
	//set the linked lists' tail to nextNode
	ll.Tail = nextNode
	return ll
}

//first need to handle pre-exinting node's head & tail
func updateCurrentNodes(ll *LinkedList, nextNode *Node) {
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
	//if the desired node to delete is the next node, set that node's Next property
	for {
		if v == node.Next.Data {
			//check that next node is not second to last in linked list
			if node.Next.Next != nil {
				node.Next = node.Next.Next
				break
			}
			//if the node to delete is second to last, the new node is now the end of linked list
			node.Next = nil
			break
		}
		node = node.Next
	}
	return ll
}