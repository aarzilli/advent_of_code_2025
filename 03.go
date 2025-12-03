package main

import (
	. "aoc/util"
	"os"
)

func main() {
	lines := Input(os.Args[1], "\n", true)
	part1 := 0
	part2 := 0
	for i := range lines {
		v := Vatoi(Spac(lines[i], "", -1))
		n := maxjoltage(2, v)
		part1 += n
		n = maxjoltage(12, v)
		part2 += n
	}
	Sol(part1)
	Sol(part2)
}

var curbest int

func maxjoltage(rem int, v []int) int {
	curbest = 0
	maxjoltageintl(0, rem, v)
	return curbest
}

func maxjoltageintl(start, rem int, v []int) {
	if rem == 0 {
		if start > curbest {
			curbest = start
		}
		return
	}
	base := 1
	for range rem - 1 {
		base *= 10
	}
	for i := range v {
		if start+((v[i]+1)*base) < curbest {
			continue
		}
		maxjoltageintl(start+(v[i]*base), rem-1, v[i+1:])
	}
}
