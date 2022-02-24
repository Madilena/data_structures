package main

import "fmt"

//graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

//vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

//AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(k int) {
	g.vertices = append(g.vertices, &Vertex{key: k})
}


func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	fmt.Println(test)
}
