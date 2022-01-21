package main

import "fmt"

func main() {
	//variables in memory need 4 things: name, type, value, and address
	i, j := 42, 2701
	fmt.Println()
	fmt.Println("i, j := 42, 2701")
	fmt.Println("printing &i:", &i)
	fmt.Println("printing &j:", &j)
	fmt.Println()
	//read & as 'address of'

	//lets store the address of i in a variable p
	p := &i
	//p is now a pointer -> the value of p is the address of i
	fmt.Println("p := &i")
	fmt.Println("printing p:",p)

	//the star in *p acts as an operator and returns the value of what p is pointing at
	//this is also called dereferencing
	//*p is the value at that address
	fmt.Println("printing *p:",*p)
	fmt.Println()
	//if you change the value of *p you are changing the value of i
	*p = 21
	fmt.Println("*p = 21")
	fmt.Println("printing *p:",*p)
	fmt.Println("printing i:",i)

	//as a note *int is a pointer type and we say that int is the base
}
