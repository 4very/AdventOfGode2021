package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/4very/AdventOfGode2021/helper"
)

type cave struct {
	label     string
	big       bool
	isStart   bool
	isEnd     bool
	neighbors []*cave
}

func addNeighbor(c *cave, n *cave) {
	(*c).neighbors = append((*c).neighbors, n)
	(*n).neighbors = append((*n).neighbors, c)
}

func pprintCavesMap(caves map[string]*cave) {
	for _, cave := range caves {
		fmt.Printf("%s (%t): ", cave.label, cave.big)
		for _, n := range cave.neighbors {
			fmt.Printf("%s ", (*n).label)
		}
		fmt.Println()
	}
	fmt.Println()
}

func pprintCavesArray(caves []*cave) {
	for _, cave := range caves {
		fmt.Printf("%s->", (*cave).label)
	}
	fmt.Println("end")
}

func sprintCavesArray(caves []*cave) string {
	str := ""
	for _, cave := range caves {
		str += fmt.Sprintf("%s,", (*cave).label)
	}
	str += "end"
	return str
}

func inArray(val *cave, array []*cave) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}

func startCaveFinding(scave *cave, pathsMap *map[string]bool) {
	var visited []*cave
	caveFindingRecut(scave, visited, pathsMap)
}

func caveFindingRecut(cave *cave, visited []*cave, pathsMap *map[string]bool) {
	if cave.isEnd {
		(*pathsMap)[sprintCavesArray(visited)] = true
		return
	}
	visited = append(visited, cave)
	for _, n := range (*cave).neighbors {
		if !inArray(n, visited) || (*n).big {
			caveFindingRecut(n, visited, pathsMap)
		}
	}
}

func startCaveFindingPt2(scave *cave, pathsMap *map[string]bool) {
	var visited []*cave
	caveFindingRecutPt2(scave, visited, pathsMap, false)
}

func caveFindingRecutPt2(cave *cave, visited []*cave, pathsMap *map[string]bool, double bool) {
	if cave.isEnd {
		(*pathsMap)[sprintCavesArray(visited)] = true
		return
	}
	visited = append(visited, cave)
	for _, n := range (*cave).neighbors {
		if !inArray(n, visited) || (*n).big {
			caveFindingRecutPt2(n, visited, pathsMap, double)
		}
		if inArray(n, visited) && !double && !(*n).big && !(*n).isEnd && !(*n).isStart {
			caveFindingRecutPt2(n, visited, pathsMap, true)
		}
	}
}

func main() {
	args := os.Args[1:]
	data := helper.ReadFile(args[0])

	var caveMap map[string]*cave = make(map[string]*cave)

	caveMap["start"] = &cave{label: "start", isStart: true}
	caveMap["end"] = &cave{label: "end", isEnd: true}

	for _, line := range data {
		r := regexp.MustCompile(`(.*)-(.*)`)
		match := r.FindStringSubmatch(line)
		cave1 := match[1]
		cave2 := match[2]

		if _, ok := caveMap[cave1]; !ok {
			isBig := cave1 != strings.ToLower(cave1)
			caveMap[cave1] = &cave{label: cave1, big: isBig}
		}

		if _, ok := caveMap[cave2]; !ok {
			isBig := cave2 != strings.ToLower(cave2)
			caveMap[cave2] = &cave{label: cave2, big: isBig}
		}

		addNeighbor(caveMap[cave1], caveMap[cave2])
	}

	pprintCavesMap(caveMap)

	var pathsMap map[string]bool = make(map[string]bool)
	startCaveFinding(caveMap["start"], &pathsMap)
	fmt.Println("Part1:", len(pathsMap))

	var pathsMap2 map[string]bool = make(map[string]bool)
	startCaveFindingPt2(caveMap["start"], &pathsMap2)
	fmt.Println("Part2:", len(pathsMap2))

}
