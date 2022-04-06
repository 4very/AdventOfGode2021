package main

import (
	"fmt"
	"math"
)

type chiton struct {
	id        int
	risk      int
	neighbors []*chiton
}

func (c *chiton) addNeighbor(n *chiton) {
	(*c).neighbors = append((*c).neighbors, n)
}

func (c *chiton) print() {

	fmt.Printf("%d: %d\n", c.id, c.risk)
	for _, n := range c.neighbors {
		fmt.Printf("  %d: %d\n", (*n).id, (*n).risk)
	}
	fmt.Println()
}

func pprint(c []*chiton) {
	for _, cn := range c {
		cn.print()
	}
}

func pprintPath(path []*chiton) {
	for _, cn := range path {
		fmt.Printf("%d ", cn.id)
	}
	fmt.Println()
}

func part1(data []string) int {

	height := len(data)
	width := len(data[0])

	var chitons []*chiton

	for y, line := range data {
		for x, c := range line {

			chiton := &chiton{
				id:        y*width + x,
				risk:      int(c - 48),
				neighbors: []*chiton{},
			}
			chitons = append(chitons, chiton)
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			chiton := chitons[y*width+x]
			if y > 0 {
				chiton.addNeighbor(chitons[(y-1)*width+x])
			}
			if y < height-1 {
				chiton.addNeighbor(chitons[(y+1)*width+x])
			}
			if x > 0 {
				chiton.addNeighbor(chitons[y*width+x-1])
			}
			if x < width-1 {
				chiton.addNeighbor(chitons[y*width+x+1])
			}
		}
	}

	return lowestPath(chitons[0], chitons[len(chitons)-1], chitons)

}

func lowestPath(start *chiton, end *chiton, graph []*chiton) int {
	var dist map[*chiton]int = make(map[*chiton]int)
	var prev map[*chiton]*chiton = make(map[*chiton]*chiton)
	var q []*chiton = []*chiton{start}

	for _, c := range graph {
		dist[c] = math.MaxInt
		prev[c] = nil
	}

	dist[start] = 0
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, n := range u.neighbors {
			alt := dist[u] + n.risk
			if alt < dist[n] {
				dist[n] = alt
				prev[n] = u
				q = append(q, n)
			}
		}
	}

	return dist[end]

}
