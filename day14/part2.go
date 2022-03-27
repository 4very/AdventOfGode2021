// time to OPTIMIZE
package main

import (
	"fmt"
	"regexp"
)

func countMap(str map[string]int) map[string]int {

	count := make(map[string]int)
	for key, val := range str {
		// fmt.Printf("%s, %d: %v\n", key, val, count)
		count[string(key[0])] += val
		count[string(key[1])] += val
		// fmt.Printf("after: %v\n", count)
	}
	return count

}

func maxMinusMinMap(str map[string]int) int {

	count := countMap(str)
	// fmt.Println(count)

	min := 999999
	max := 0
	for _, v := range count {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min

}

var ans []string = []string{
	"NNCB",
	"NCNBCHB",
	"NBCCNBBBCBHCB",
	"NBBBCNCCNBBNBNBBCHBHHBCHB",
	"NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB",
}

func part2(data []string) int {

	var str map[string]int = make(map[string]int)

	var instructions map[string]string = make(map[string]string)
	instrRegex := regexp.MustCompile(`(.*) -> (.*)`)

	for _, v := range data[2:] {
		match := instrRegex.FindStringSubmatch(v)
		instructions[match[1]] = match[2]
	}

	for i := 1; i < len(data[0]); i++ {
		str[data[0][i-1:i+1]]++
	}

	for step := 0; step < 4; step++ {

		var newstr map[string]int = make(map[string]int)

		for key, val := range str {
			if _, ok := instructions[key]; ok {
				left := string(key[0]) + instructions[key]
				right := instructions[key] + string(key[1])
				newstr[left] += val
				newstr[right] += val

			} else {
				newstr[key] += val
			}
		}
		str = newstr
		fmt.Printf("%d: %v\n", step+1, str)
		fmt.Printf("%d: %v\n", step+1, split(ans[step+1]))
	}

	return maxMinusMinMap(str)

}
