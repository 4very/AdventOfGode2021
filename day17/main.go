package main

import (
	"fmt"
	"os"
	"time"

	"github.com/4very/AdventOfGode2021/helper"
)

func main() {
	args := os.Args[1:]
	data := helper.ReadFile(args[0])

	start := time.Now()
	fmt.Println("part1:", part1(data))
	fmt.Println("In:", time.Since(start))
	start = time.Now()
	fmt.Println("part2:", part2(data))
	fmt.Println("In:", time.Since(start))

}
