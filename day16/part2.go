package main

import "fmt"

func parsePt2(b *[]bool) uint64 {
	msbpi(b, 3)
	typeid := msbpi(b, 3)
	switch typeid {
	case 0:
		return t0(b)
	case 1:
		return t1(b)
	case 2:
		return t2(b)
	case 3:
		return t3(b)
	case 4:
		return t4(b)
	case 5:
		return t5(b)
	case 6:
		return t6(b)
	case 7:
		return t7(b)
	}

	return 0
}

func t0(b *[]bool) uint64 {
	vals := tn4Pt2(b)
	ret := uint64(0)
	for _, val := range vals {
		ret += val
	}
	return ret
}

func t1(b *[]bool) uint64 {
	vals := tn4Pt2(b)
	ret := uint64(1)
	for _, val := range vals {
		ret *= val
	}
	return ret
}

func t2(b *[]bool) uint64 {
	vals := tn4Pt2(b)
	ret := ^uint64(0)
	for _, val := range vals {
		if val < ret {
			ret = val
		}
	}
	return ret
}

func t3(b *[]bool) uint64 {
	vals := tn4Pt2(b)
	ret := uint64(0)
	for _, val := range vals {
		if val > ret {
			ret = val
		}
	}
	return ret
}

func t5(b *[]bool) uint64 {
	vals := tn4Pt2(b)
	if vals[0] > vals[1] {
		return 1
	}
	return 0
}

func t6(b *[]bool) uint64 {
	vals := tn4Pt2(b)
	if vals[0] < vals[1] {
		return 1
	}
	return 0
}

func t7(b *[]bool) uint64 {
	vals := tn4Pt2(b)
	if vals[0] == vals[1] {
		return 1
	}
	return 0
}

func tn4Pt2(b *[]bool) []uint64 {
	ltypeid := msbp(b, 1)

	if ltypeid[0] {
		return tn4l1Pt2(b)
	} else {
		return tn4l0Pt2(b)
	}
}

func tn4l0Pt2(b *[]bool) []uint64 {
	bytes := msbpi(b, 15)
	subpacket := msbp(b, bytes)

	if debug {
		fmt.Printf("%db\n-----\n", bytes)
	}
	var ret []uint64 = make([]uint64, 0)

	for {
		ret = append(ret, parsePt2(&subpacket))
		if len(subpacket) <= 0 {
			break
		}
	}
	if debug {
		fmt.Printf("-----\n")
	}
	return ret
}

func tn4l1Pt2(b *[]bool) []uint64 {
	packets := msbpi(b, 11)
	var ret []uint64 = make([]uint64, 0)

	if debug {
		fmt.Printf("%dp\n-----\n", packets)
	}

	for i := uint64(0); i < packets; i++ {
		ret = append(ret, parsePt2(b))
	}

	if debug {
		fmt.Printf("-----\n")
	}

	return ret
}

func part2(data []string) uint64 {
	barray := newPacket(data[0])
	// print(barray)
	return parsePt2(&barray)
}
