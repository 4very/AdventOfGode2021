package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// get biggest x velocity where xmin is the sum of all
// natural numbers less than it

func sigmaInverse(i float64) float64 {
	return (math.Sqrt(8*float64(i)+1) - 1) / 2
}

func sigma(i float64) float64 {
	return (float64(i) * (float64(i) + 1)) / 2
}

func maxX(xmin int64) int64 {
	xminf := float64(xmin)
	return int64(math.Ceil(sigmaInverse(xminf)))
}

func maxY(ymax int64) int64 {
	ymaxf := float64(ymax)
	return int64(math.Abs(ymaxf) - 1)

}

type input struct {
	xmin, xmax, ymin, ymax int64
}

func parseInput(s string) input {
	re := regexp.MustCompile(`target area: x=(-?\d*)..(-?\d*), y=(-?\d*)..(-?\d*)`)
	vals := re.FindStringSubmatch(s)

	var ret input

	ret.xmin, _ = strconv.ParseInt(vals[1], 10, 64)
	ret.xmax, _ = strconv.ParseInt(vals[2], 10, 64)
	ret.ymin, _ = strconv.ParseInt(vals[3], 10, 64)
	ret.ymax, _ = strconv.ParseInt(vals[4], 10, 64)

	return ret
}

func part1(data []string) string {

	pin := parseInput(data[0])

	return fmt.Sprintf("velocity (%.0d,%.0d) with max height %.0f", maxX(pin.xmin), maxY(pin.ymin), sigma(float64(maxY(pin.ymin))))
}
