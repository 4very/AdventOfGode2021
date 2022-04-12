package main

import (
	"fmt"
	"math"
)

var sample map[[2]int64]bool = map[[2]int64]bool{
	{23, -10}: true,
	{25, -9}:  true,
	{27, -5}:  true,
	{29, -6}:  true,
	{22, -6}:  true,
	{21, -7}:  true,
	{9, 0}:    true,
	{27, -7}:  true,
	{24, -5}:  true,
	{25, -7}:  true,
	{26, -6}:  true,
	{25, -5}:  true,
	{6, 8}:    true,
	{11, -2}:  true,
	{20, -5}:  true,
	{29, -10}: true,
	{6, 3}:    true,
	{28, -7}:  true,
	{8, 0}:    true,
	{30, -6}:  true,
	{29, -8}:  true,
	{20, -10}: true,
	{6, 7}:    true,
	{6, 4}:    true,
	{6, 1}:    true,
	{14, -4}:  true,
	{21, -6}:  true,
	{26, -10}: true,
	{7, -1}:   true,
	{7, 7}:    true,
	{8, -1}:   true,
	{21, -9}:  true,
	{6, 2}:    true,
	{20, -7}:  true,
	{30, -10}: true,
	{14, -3}:  true,
	{20, -8}:  true,
	{13, -2}:  true,
	{7, 3}:    true,
	{28, -8}:  true,
	{29, -9}:  true,
	{15, -3}:  true,
	{22, -5}:  true,
	{26, -8}:  true,
	{25, -8}:  true,
	{25, -6}:  true,
	{15, -4}:  true,
	{9, -2}:   true,
	{15, -2}:  true,
	{12, -2}:  true,
	{28, -9}:  true,
	{12, -3}:  true,
	{24, -6}:  true,
	{23, -7}:  true,
	{25, -10}: true,
	{7, 8}:    true,
	{11, -3}:  true,
	{26, -7}:  true,
	{7, 1}:    true,
	{23, -9}:  true,
	{6, 0}:    true,
	{22, -10}: true,
	{27, -6}:  true,
	{8, 1}:    true,
	{22, -8}:  true,
	{13, -4}:  true,
	{7, 6}:    true,
	{28, -6}:  true,
	{11, -4}:  true,
	{12, -4}:  true,
	{26, -9}:  true,
	{7, 4}:    true,
	{24, -10}: true,
	{23, -8}:  true,
	{30, -8}:  true,
	{7, 0}:    true,
	{9, -1}:   true,
	{10, -1}:  true,
	{26, -5}:  true,
	{22, -9}:  true,
	{6, 5}:    true,
	{7, 5}:    true,
	{23, -6}:  true,
	{28, -10}: true,
	{10, -2}:  true,
	{11, -1}:  true,
	{20, -9}:  true,
	{14, -2}:  true,
	{29, -7}:  true,
	{13, -3}:  true,
	{23, -5}:  true,
	{24, -8}:  true,
	{27, -9}:  true,
	{30, -7}:  true,
	{28, -5}:  true,
	{21, -10}: true,
	{7, 9}:    true,
	{6, 6}:    true,
	{21, -5}:  true,
	{27, -10}: true,
	{7, 2}:    true,
	{30, -9}:  true,
	{21, -8}:  true,
	{22, -7}:  true,
	{24, -9}:  true,
	{20, -6}:  true,
	{6, 9}:    true,
	{29, -5}:  true,
	{8, -2}:   true,
	{27, -8}:  true,
	{30, -5}:  true,
	{24, -7}:  true,
}

func xminvel(xval int64) int64 {
	return int64(math.Ceil(sigmaInverse(float64(xval))))
}

func xt(xmin int64, xmax int64, xvel int64, xtmax int64) [2]int64 {
	var xtmin, x, t int64 = 0, 0, 0

	for xval := xvel; xval > 0; xval-- {
		x += xval
		t++
		if x >= xmin && xtmin == 0 && x <= xmax {
			xtmin = t
		}
		if x > xmax {
			xtmax = t
			break
		}
	}
	ret := [2]int64{xtmin, xtmax}

	return ret
}

func xvels(xmin int64, xmax int64, maxx int64) map[int64][]int64 {
	ret := make(map[int64][]int64)
	xrange := xmax - xmin

	// sigma-1(xmin) ... xrange
	// we know for sure
	for i := xminvel(xmin); i <= xrange; i++ {
		time := xt(xmin, xmax, i, maxx)
		for j := time[0]; j < time[1]; j++ {
			ret[j] = append(ret[j], i)
		}
	}

	// xrange ... xmin
	// we dont know for sure
	for i := xrange + 1; i < xmin; i++ {
		time := xt(xmin, xmax, i, maxx)
		if time[0] != 0 {
			for j := time[0]; j < time[1]; j++ {
				ret[j] = append(ret[j], i)
			}
		}
	}

	// xmin ... xmax
	// we know for sure
	for i := xmin; i <= xmax; i++ {
		ret[1] = append(ret[1], i)
	}

	return ret
}

func yt(ymin int64, ymax int64, yvi int64) [2]int64 {
	var ytmax, ytmin, y, t int64 = 0, 0, 0, 0
	yv := yvi

	for {
		y += yv
		t++
		yv--

		if y <= ymax && ytmin == 0 && y >= ymin {
			ytmin = t
		}
		if y < ymin {
			ytmax = t
			break
		}
	}
	ret := [2]int64{ytmin, ytmax}

	return ret
}

func abs(i int64) int64 {
	return int64(math.Abs(float64(i)))
}

func yvels(ymin int64, ymax int64) map[int64][]int64 {

	ret := make(map[int64][]int64)
	// yrange := abs(ymax - ymin)

	// ymin .. ymax
	// we know t = 1
	for y := ymin; y <= ymax; y++ {
		ret[1] = append(ret[1], y)
	}

	// ymax .. abs(ymax)
	// we dont know
	for y := ymax; y < abs(ymax); y++ {
		time := yt(ymin, ymax, y)
		for i := time[0]; i < time[1]; i++ {
			ret[i] = append(ret[i], y)
		}
	}

	// abs(max) .. abs(ymin)-1
	// we know t = 2v+1
	for y := abs(ymax); y < abs(ymin); y++ {
		ret[2*y+2] = append(ret[2*y+2], y)

	}

	return ret
}

func countMap(m map[[2]int64]bool) int {
	ret := 0
	for _, v := range m {
		if v {
			ret++
		}
	}
	return ret
}

func pprint(m map[[2]int64]bool) {
	for key := range m {
		fmt.Printf("(%d, %d)\n", key[0], key[1])
	}
}

func compare(m1 map[[2]int64]bool, m2 map[[2]int64]bool) {
	for key, val := range m1 {
		if val {
			if !m2[key] {
				fmt.Printf("In m1 but not m2: (%d, %d)\n", key[0], key[1])
			} else {
				m2[key] = false
			}
		}
	}

	for key, val := range m2 {
		if val {
			fmt.Printf("In m2 but not in m1: (%d, %d)\n", key[0], key[1])
		}
	}
}

func part2(data []string) int {
	pin := parseInput(data[0])

	velocities := make(map[[2]int64]bool)

	xvs := xvels(pin.xmin, pin.xmax, 2*abs(pin.ymin)+1)
	yvs := yvels(pin.ymin, pin.ymax)

	for i := int64(1); i <= 2*abs(pin.ymin); i++ {
		for _, xval := range xvs[i] {
			for _, yval := range yvs[i] {
				velocities[[2]int64{xval, yval}] = true
			}
		}
	}
	return countMap(velocities)
}
