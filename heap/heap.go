package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{100, 3, 25, 29, 52, 1, 8, 9, 11, 18, 19, 69, 75, 92, 22}
	array_length := len(array)
	//startIndex is last non-leaf node
	startIndex := array_length/2 - 1
	fmt.Println(array)
	//sift down from last non-leaf node
	for i := startIndex; i >= 0; i-- {
		heapify(&array, i)
	}
	fmt.Println("number of tree levels: ", levelOrder(len(array)))

	heapAsTreeBasic(&array)
	printHeapAsTree(&array)
}

func heapify(heap *[]int, i int) {
	smallest := i
	lChildIndex := leftChild(i)
	rChildIndex := rightChild(i)

	if childIsSmaller(lChildIndex, heap, smallest) {
		smallest = lChildIndex
	}

	if childIsSmaller(rChildIndex, heap, smallest) {
		smallest = rChildIndex
	}

	if smallest != i {
		swapCurrentSmallestWithChild(heap, i, smallest)
		heapify(heap, smallest)
	}
}

func heapAsTree(heap *[]int) [][]int {
	order := levelOrder(len(*heap))
	orderLevelArray := [][]int{}
	for i := 0; i <= order; i++ {
		bushArray := []int{}
		if i == 0 {
			continue
		}
		for j := powerOfTwo(i-1) - 1; j < powerOfTwo(i)-1; j++ {
			if j < len(*heap) {
				bushArray = append(bushArray, (*heap)[j])
			}
		}
		orderLevelArray = append(orderLevelArray, bushArray)
	}
	//fmt.Println("this is the order array:", orderLevelArray)
	return orderLevelArray
}

func printHeapAsTree(heap *[]int) {
	order := levelOrder(len(*heap))
	for i := 0; i < order; i++ {
		fmt.Println(heapAsTree(heap)[i])
	}
}

func heapAsTreeBasic(heap *[]int) {
	order := levelOrder(len(*heap))
	for i := 0; i <= order; i++ {
		if i == 0 {
			continue
		}
		for j := powerOfTwo(i-1) - 1; j < powerOfTwo(i)-1; j++ {
			if j < len(*heap) {
				fmt.Println("tree level is ", i, "array index is: ", j, "node value is ", (*heap)[j], " ")
			}
		}
	}
}

func levelOrder(arrayLength int) int {
	binLog := math.Log2(float64(arrayLength))
	roundUp := math.Ceil(binLog)
	return int(roundUp)
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func childIsSmaller(child int, heap *[]int, smallest int) bool {
	return childIsWithinArray(child, heap) && childIsLessThanSmallest(child, heap, smallest)
}

func childIsWithinArray(child int, heap *[]int) bool {
	return child < len(*heap)
}

func childIsLessThanSmallest(child int, heap *[]int, smallest int) bool {
	return (*heap)[child] < (*heap)[smallest]
}

func swapCurrentSmallestWithChild(heap *[]int, i int, smallest int) {
	(*heap)[i], (*heap)[smallest] = (*heap)[smallest], (*heap)[i]
	fmt.Println((*heap))
}

func powerOfTwo(exponent int) int {
	return int(math.Pow(2, float64(exponent)))
}

func depthOfTree(exponent int) int {
	return int(math.Pow(2, float64(exponent)))
}
