package main

import "fmt"

func pprintMap(chitons []*chiton, width int) {
	for i, c := range chitons {
		if i%width == 0 {
			fmt.Println("---")
		}
		fmt.Printf("%d", c.risk)
	}

	fmt.Println()
}

func part2(data []string) int {
	height := len(data) * 5
	width := len(data[0]) * 5

	var chitons []*chiton

	for i := 0; i < 5; i++ {
		for y, line := range data {
			for j := 0; j < 5; j++ {
				for x, c := range line {

					chiton := &chiton{
						id:        y*width + x,
						risk:      (int(c-48)+i+j-1)%9 + 1,
						neighbors: []*chiton{},
					}
					chitons = append(chitons, chiton)
				}
			}
		}
	}

	// pprintMap(chitons[:50], 10)

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
	path := lowestPath(chitons[0], chitons[len(chitons)-1], chitons)
	return path

}
