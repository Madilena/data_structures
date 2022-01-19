package main

import "fmt"

//ArraySize is the size of the hash table
const ArraySize = 7

//hash table will hold an array of buckets
type HashTable struct {
	array [ArraySize]*bucket
}

//bucket is a linked list in each slot of the array
type bucket struct {
	head *bucketNode
}

//bucketNode is a linked list node that holds reference to next node in list
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

//Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

//Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

//insert will take in a key, create a node with the key and insert the node into the bucket
func (b *bucket) insert(k string) {
	//if key is not already in bucket
	if !b.search(k) {
		//create bucketNode with key
		//new node we want to insert will be the head
		newNode := &bucketNode{key: k}
		//set new node as head and set the next of new node to the head that was there before
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, " already exists")
	}
}

//search will take in a key and return true if the bucket has the key
//keep looping until we find a match - so we need to loop until the current node is empty
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		//if no match then we update the currentNode to next node
		currentNode = currentNode.next
	}
	return false
}

//delete and search are very similar except in how they define the currentNode

//delete will take in a key and delete the node from the bucket
//linked list will delete the current node by skipping the current node
//and making the previous node point to the next node
//and we cant do that if we go actually over the previous node
//so we're going to put each node into a previous node and express current node
//as previousNode.next
func (b *bucket) delete(k string) {

	//if the key we want to delete is the head we set the head to next node and come out of the method
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
			return
		}
		//if no match then we update the currentNode to next node
		previousNode = previousNode.next
	}
}

//hash function
func hash(key string) int {
	sum := 0
	//loop through each character in string and add value to sum
	for _, v := range key {
		sum += int(v)
	}
	fmt.Println(key, "has ASCII code sum:", sum, " and will be at index ", sum%ArraySize)
	return sum % ArraySize
}

//Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"WATSON", "SAMSON", "WENDY", "MADILENA", "BAGEL", "MAPOTOFU",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("BLUEDRINK")
	hashTable.Delete("BAGEL")
	fmt.Println(hashTable)
}
