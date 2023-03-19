package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello")

	graph := Graph{}

	vertices := []string{"A", "B", "C", "D", "E"}
	graph.AddVetices(vertices)
	graph.AddEdges("A", "B", -1)
	graph.AddEdges("A", "C", 4)
	graph.AddEdges("B", "C", 3)
	graph.AddEdges("B", "E", 2)
	graph.AddEdges("B", "D", 2)
	graph.AddEdges("D", "B", 1)
	graph.AddEdges("D", "C", 5)
	graph.AddEdges("E", "D", -3)

	graph.Print()
	graph.ApplyBellmanFord("A")
	graph.PrintDistances()

}
