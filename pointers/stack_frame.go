package main

import "fmt"

func main() {
	a := 4
	//this is only making changes inside this active frame. does not change anything else.
	//not efficent bc we are copying the arguments each time we make a function call
	//when function call ends and active frame goes back to main() function, a is still 4
	squareValue(a)

	//we are copying the address of a and assigning as a pointer
	//and that pointer is pointing across the stack frame 
	//and this is how we can actulaly change a
	squareAddress(&a)
}

//immutability
func squareValue(v int){
	v *= v
	fmt.Println(&v, v)
}

//efficiency
func squareAddress(p *int){
	*p *= *p
	fmt.Println(p, *p)
}
