package main

import (
	"fmt"
	"os"

	"github.com/4very/AdventOfGode2021/helper"
)

func split(s string) map[string]int {
	returnval := make(map[string]int)
	for i := 1; i < len(s); i++ {
		returnval[s[i-1:i+1]]++
	}
	return returnval
}

func main() {
	args := os.Args[1:]
	data := helper.ReadFile(args[0])
	fmt.Println("part1:", part1(data))
	fmt.Println("part2:", part2(data))
	fmt.Println()
	fmt.Println("example:", split("NBBBCNCCNBBNBNBBCHBHHBCHB"))

}
