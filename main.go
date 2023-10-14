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

	array := []interface{}{
		Edge{U: "A", V: "B", Weight: 1},
		Edge{U: "A", V: "C", Weight: 3},
		Edge{U: "B", V: "C", Weight: 2},
		Edge{U: "C", V: "D", Weight: 5},
		Edge{U: "D", V: "E", Weight: -1},
		Edge{U: "D", V: "F", Weight: 6},
	}

	fmt.Println("Mergesort...")

	s := MergeSort(array, 0, len(array)-1, func(i, j int) bool {
		left := array[i].(Edge)
		right := array[j].(Edge)
		return left.Weight > right.Weight
	})

	for _, e := range s {
		fmt.Println(e.(Edge).Weight)
	}

	fmt.Println("Ouicksort on same array...")

	qs := QuickSort(array, 0, len(array)-1, func(i, j int) bool {
		left := array[i].(Edge)
		right := array[j].(Edge)
		return left.Weight > right.Weight
	})

	for _, e := range qs {
		fmt.Println(e.(Edge).Weight)
	}

}
