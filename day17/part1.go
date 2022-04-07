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

func maxX(xmin float64) float64 {
	return math.Ceil(sigmaInverse(xmin))
}

func maxY(ymax float64) float64 {
	return math.Abs(ymax) - 1

}

func part1(data []string) string {
	re := regexp.MustCompile(`target area: x=(-?\d*)..(-?\d*), y=(-?\d*)..(-?\d*)`)
	vals := re.FindStringSubmatch(data[0])
	xmin, _ := strconv.ParseFloat(vals[1], 64)
	ymax, _ := strconv.ParseFloat(vals[3], 64)

	return fmt.Sprintf("velocity (%.0f,%.0f) with max height %.0f", maxX(xmin), maxY(ymax), sigma(maxY(ymax)))
}
