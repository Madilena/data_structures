package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{100, 3, 25, 1}
	array_length := len(array)
	for i := 0; i < array_length; i++ {
		heapify(&array, i)
	}
	fmt.Println(array)

}

func powerOfTwo(exponent int) int {
	return int(math.Pow(2, float64(exponent)))
}

func depthOfTree(exponent int) int {
	return int(math.Pow(2, float64(exponent)))
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func heapify(heap *[]int, i int) {
	smallest := i
	lChild := leftChild(i)
	rChild := rightChild(i)
	if lChild < len(*heap) && (*heap)[lChild] < (*heap)[smallest] {
		smallest = lChild
	}

	if rChild < len(*heap) && (*heap)[rChild] < (*heap)[smallest] {
		smallest = rChild
	}
	if smallest != i {
		(*heap)[i], (*heap)[smallest] = (*heap)[smallest], (*heap)[i]
		heapify(heap, smallest)
	}
}
