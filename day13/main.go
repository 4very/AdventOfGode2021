package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/4very/AdventOfGode2021/helper"
)

type paper struct {
	height, width int
	points        map[[2]int]bool
}

func (paper *paper) addPoint(x, y int) {
	(*paper).points[[2]int{x, y}] = true
	(*paper).h(y)
	(*paper).w(x)
}

func (paper *paper) h(y int) {
	if y > (*paper).height {
		(*paper).height = y
	}
}

func (paper *paper) w(x int) {
	if x > (*paper).width {
		(*paper).width = x
	}
}

func (p *paper) print() {
	for y := 0; y <= (*p).height; y++ {
		for x := 0; x <= (*p).width; x++ {
			if (*p).points[[2]int{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (p *paper) fold(folds [2]int) {
	dir := folds[0]
	where := folds[1]
	var newkey [2]int

	for key := range (*p).points {
		if dir == 0 && key[0] > where {
			newkey = [2]int{(*p).width - key[0], key[1]}
		} else if dir == 1 && key[1] > where {
			newkey = [2]int{key[0], (*p).height - key[1]}
		} else {
			continue
		}
		(*p).points[newkey] = true
		delete((*p).points, key)
	}

	if dir == 0 {
		(*p).width = where - 1
	} else {
		(*p).height = where - 1
	}
}

func (p *paper) pcount() {
	count := 0
	for _, v := range (*p).points {
		if v {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	args := os.Args[1:]
	data := helper.ReadFile(args[0])

	var p *paper = &paper{points: make(map[[2]int]bool)}
	var folds [][2]int // 0 if x direction, 1 if y direction

	cordsreg := regexp.MustCompile(`(\d*),(\d*)`)
	foldsreg := regexp.MustCompile(`fold along (.)\=(\d*)`)

	for _, line := range data {
		if strings.Contains(line, ",") {
			match := cordsreg.FindStringSubmatch(line)
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			p.addPoint(x, y)
		}

		if strings.Contains(line, "fold") {
			match := foldsreg.FindStringSubmatch(line)
			dirbool := match[1] == "y"
			dir := 0
			if dirbool {
				dir = 1
			}
			val, _ := strconv.Atoi(match[2])
			folds = append(folds, [2]int{dir, val})
		}
	}

	for _, fold := range folds {
		p.fold(fold)
	}
	p.print()
	p.pcount()
}
