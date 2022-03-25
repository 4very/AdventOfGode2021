package main

import (
	"fmt"
	"os"

	"github.com/4very/AdventOfGode2021/helper"
)

type Octint interface {
	isMax() bool      // is the oct max energy
	flash() bool      // flash the oct
	addNeighbor(*Oct) // add a neighbor
}

type Oct struct {
	x, y      int
	neighbors []*Oct
	power     int
}

func (o Oct) isMax() bool {
	return o.power > 9
}

// increases the power of neighboring octs
// if the power is 9, return
func (o Oct) flash() []*Oct {
	var result []*Oct
	for _, oct := range o.neighbors {
		(*oct).power++
		if (*oct).isMax() {
			result = append(result, oct)
		}
	}
	return result
}

// TODO
func addNeighbor(o *Oct, op *Oct) {
	(*o).neighbors = append((*o).neighbors, op)
}

func increaseAll(octs *[]Oct) {

	for _, oct := range *octs {
		oct.power++
	}

}

func inArray(val *Oct, array []*Oct) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}

func pprintoctGrid(octs []Oct) {
	for i, oct := range octs {
		fmt.Printf("%d", (oct).power)
		if i%10 == 9 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func pprintoctLong(octs []*Oct) {
	for _, oct := range octs {
		fmt.Printf("(%d,%d): %d\n", (*oct).x, (*oct).y, (*oct).power)
	}
	fmt.Println()
}

var mods [8][2]int = [8][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func main() {
	args := os.Args[1:]
	data := helper.ReadFile(args[0])

	// create octs
	var octs []Oct
	var octMap = make(map[[2]int]Oct)
	for i, line := range data {
		for j, pow := range line {
			octMap[[2]int{j, i}] = Oct{x: j, y: i, power: int(pow) - 48}
			octs = append(octs, octMap[[2]int{j, i}])
		}
	}

	var count int = 0

	for i := 0; i < 100; i++ {

		for _, mod := range mods {
			xmod := i%10 + mod[0]
			ymod := int(i/10) + mod[1]

			if xmod < 0 || xmod > 9 || ymod < 0 || ymod > 9 {
				continue
			}

			addNeighbor(&octs[i], &octs[ymod*10+xmod])
		}
	}

	fmt.Println("Starting grid:")
	pprintoctGrid(octs)

	for step := 0; step < 400; step++ {

		var queue []*Oct = []*Oct{}
		var flashed []*Oct = []*Oct{}

		for i := range octs {
			octs[i].power++
			if octs[i].isMax() {
				queue = append(queue, &octs[i])
				// pprintoctLong(queue)
			}
		}

		// pprintoctLong(queue)

		for len(queue) > 0 {
			curoct := queue[0]
			flashed = append(flashed, curoct)
			queue = queue[1:]
			for _, oct := range curoct.flash() {
				if !inArray(oct, queue) && !inArray(oct, flashed) {
					queue = append(queue, oct)
				}
			}
		}

		count += len(flashed)

		// part 2
		if len(flashed) == 100 {
			fmt.Println("Simultaneous flash:", step+1)
		}

		if step == 99 {
			fmt.Println("After 100 steps:", count)
		}

		for _, oct := range flashed {
			oct.power = 0
		}
	}
	pprintoctGrid(octs)
	fmt.Println("Count:", count)
}
