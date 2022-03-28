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
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

//AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int){
	// get addresses of vertices
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil {
		err:= fmt.Errorf("invalid edge (%v --> %v)",from, to)
		fmt.Println(err.Error())
	}else if (contains(fromVertex.adjacent, to)){
		err:= fmt.Errorf("existing edge (%v --> %v)",from, to)
		fmt.Println(err.Error())
	}else{
	// add edge
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}}

func (g *Graph) getVertex(k int) *Vertex{
	for i, v := range g.vertices{
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

//contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

//Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertext %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
}

func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(0,1)
	test.AddEdge(1,2)
	test.AddEdge(6,2)
	test.AddEdge(3,2)
	test.AddEdge(1,2)
	test.AddEdge(5,1)
	test.Print()
}
