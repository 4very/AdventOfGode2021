package main

import (
	"fmt"
	"regexp"
)

func count(s string) map[rune]int {

	count := make(map[rune]int)
	for _, v := range s {
		count[v]++
	}
	return count

}

func maxMinusMin(s string) int {
	c := count(s)
	min := 999999
	max := 0
	for _, v := range c {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

func part1(data []string) int {

	var str string = data[0]

	fmt.Println(str)

	var instructions map[string]string = make(map[string]string)
	instrRegex := regexp.MustCompile(`(.*) -> (.*)`)

	for _, v := range data[2:] {
		match := instrRegex.FindStringSubmatch(v)
		instructions[match[1]] = match[2]
	}

	for step := 0; step < 10; step++ {
		var newstr string = ""

		for i := 1; i < len(str); i++ {
			pair := str[i-1 : i+1]
			if _, ok := instructions[pair]; ok {
				newstr += string(pair[0]) + instructions[pair]
			} else {
				newstr += string(pair[1])
			}
		}
		newstr += string(str[len(str)-1])
		str = newstr

	}

	return maxMinusMin(str)

}
