package main

import (
	"fmt"
	"strconv"
)

var debug bool = false

type literal struct {
	val uint64
	ver uint64
}

func newPacket(hex string) []bool {
	var p []bool = make([]bool, 0, len(hex)*4)
	// this goes rlly fucking fast
	for _, char := range hex {
		parse, _ := strconv.ParseUint(string(char), 16, 4)

		for _, char := range fmt.Sprintf("%04b", parse) {
			if char == '0' {
				p = append(p, false)
			}
			if char == '1' {
				p = append(p, true)
			}
		}
	}
	return p
}

// speedy
func toint(b []bool) uint64 {
	val := uint64(0)
	for _, bit := range b {
		bitval := uint64(0)
		if bit {
			bitval = 1
		}
		val = (val << 1) | bitval
	}
	return val
}

func msbi(b *[]bool, len uint64) uint64 {
	return toint(msb(b, len))
}

func msb(b *[]bool, len uint64) []bool {
	return (*b)[:len]
}

func msbp(b *[]bool, len uint64) []bool {
	ret := msb(b, len)
	(*b) = (*b)[len:]
	return ret
}

func msbpi(b *[]bool, len uint64) uint64 {
	ret := msbi(b, len)
	(*b) = (*b)[len:]
	return ret
}

func parse(b *[]bool) []literal {
	version := msbpi(b, 3)
	typeid := msbpi(b, 3)

	if typeid == 4 {
		if debug {
			fmt.Printf("lit (%d): ", version)
		}
		return []literal{{t4(b), version}}

	} else {
		if debug {
			fmt.Printf("ope (%d): ", version)
		}
		return append(tn4(b), literal{0, version})
	}
}

func t4(b *[]bool) uint64 {
	var ret uint64 = uint64(0)
	for {
		last := msbpi(b, 1)
		val := msbpi(b, 4)
		ret = ret<<4 | val

		if last == 0 {
			if debug {
				fmt.Printf("%04b = %d\n", val, ret)
			}
			return ret
		}

		if debug {
			fmt.Printf("%04b ", val)
		}
	}
}

func tn4(b *[]bool) []literal {
	ltypeid := msbp(b, 1)

	if ltypeid[0] {
		return tn4l1(b)
	} else {
		return tn4l0(b)
	}
}

func tn4l0(b *[]bool) []literal {
	bytes := msbpi(b, 15)
	subpacket := msbp(b, bytes)

	if debug {
		fmt.Printf("%db\n-----\n", bytes)
	}
	var ret []literal = make([]literal, 0)

	for {
		ret = append(ret, parse(&subpacket)...)
		if len(subpacket) <= 0 {
			break
		}
	}
	if debug {
		fmt.Printf("-----\n")
	}
	return ret
}

func tn4l1(b *[]bool) []literal {
	packets := msbpi(b, 11)
	var ret []literal = make([]literal, 0)

	if debug {
		fmt.Printf("%dp\n-----\n", packets)
	}

	for i := uint64(0); i < packets; i++ {
		ret = append(ret, parse(b)...)
	}

	if debug {
		fmt.Printf("-----\n")
	}

	return ret
}

func print(b []bool) {
	for _, bval := range b {
		val := int8(0)
		if bval {
			val = 1
		}
		fmt.Printf("%d", val)
	}
	fmt.Println()
}

func sumver(l []literal) uint64 {
	ret := uint64(0)
	for _, lit := range l {
		ret += lit.ver
	}
	return ret
}

// hex: D2FE28
// binary: 110100101111111000101000

func part1(data []string) uint64 {

	barray := newPacket(data[0])
	// print(barray)
	vals := parse(&barray)
	return sumver(vals)
}
