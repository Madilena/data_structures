package main

import (
	"fmt"
)
func main() {
	array := []int{100, 3, 25, 29, 52, 1, 8, 9, 11, 18, 19, 69, 75, 92, 22}
	printElementInArray(&array, 0)
	printAllElementsInArray(&array, len(array))
}

//O(1) -> Constant Time
//The array could have size=1 or 10000 but the function would still require 1 step
func printElementInArray(arr *[]int, i int) {
	fmt.Println((*arr)[i])
}

//O(n) -> Linear Time
//If n = 10 , we have 10 operations.  If n=1000, we have 1000 operations.  Think slope = 1
func printAllElementsInArray(arr *[]int, size int) {
	for i := 0; i < size; i++ {
		fmt.Println((*arr)[i])
	}
}
