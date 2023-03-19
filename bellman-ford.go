package main

import (
	"fmt"
	"math"
)

type Edge struct {
	U      string
	V      string
	Weight int
}

type Graph struct {
	V     []string
	dist  map[string]int
	Edges []Edge
}

func (g *Graph) Print() {
	fmt.Println("Printing the Graph")
	for _, e := range g.Edges {
		fmt.Printf("Edge %s-%s Weight %d \n", e.U, e.V, e.Weight)
	}
}

func (g *Graph) PrintDistances() {
	fmt.Println("Printing shortest distance from selected node")
	for v, d := range g.dist {
		fmt.Printf("Distance to %s is %d\n", v, d)
	}
}

func (g *Graph) AddVetices(v []string) {
	if len(v) == 0 {
		fmt.Println("Error No Vertices provided")
		return
	}
	g.V = v
}

func (g *Graph) AddEdges(u, v string, w int) {
	edge := Edge{
		U:      u,
		V:      v,
		Weight: w,
	}
	g.Edges = append(g.Edges, edge)
}

func (g *Graph) ApplyBellmanFord(src string) {
	g.dist = make(map[string]int)

	for _, v := range g.V {
		if v == src {
			g.dist[v] = 0
		} else {
			g.dist[v] = math.MaxInt
		}
	}

	for i := 0; i < len(g.V)-1; i += 1 {
		for _, e := range g.Edges {
			if g.dist[e.U] != math.MaxInt &&
				g.dist[e.V] > g.dist[e.U]+e.Weight {
				g.dist[e.V] = g.dist[e.U] + e.Weight
			}
		}
	}

	fmt.Println("Applied Bellman Ford")
}
