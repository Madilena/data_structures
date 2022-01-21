package main

import "fmt"

type person struct {
	name string
	age  int
}

func initPersonValue() person {
	m := person{name: "Samson", age: 1}
	fmt.Println("initPersonValue -->", m)
	return m
}

func initPersonAddress() *person {
	m := person{name: "Watson", age: 5}
	fmt.Printf("initPersonAddress --> %p\n", &m)
	return &m
}

func main() {
	//because initPerson() is returning a value, it is making a copy of that in active frame
	fmt.Println("main calling initPersonValue() --> ", initPersonValue())

	//in the stack, after initPersonAddress() call is finished, the frame is invalid
	//so the address we copied into the active frame is useless, its gone.
	//so, this is actually going to be allocated on the heap
	fmt.Printf("main calling initPersonAddress() --> %p\n", initPersonAddress())
}
