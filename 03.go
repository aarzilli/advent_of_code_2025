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
		n := maxjoltage(pow10(2), v)
		part1 += n
		n = maxjoltage(pow10(12), v)
		part2 += n
	}
	Sol(part1)
	Sol(part2)
}

var curbest int

func maxjoltage(base int, v []int) int {
	curbest = 0
	maxjoltageintl(0, base, v)
	return curbest
}

func maxjoltageintl(start, base int, v []int) {
	for i := range v {
		if start+((v[i]+1)*base) < curbest {
			continue
		}
		n := start + (v[i] * base)
		if base == 1 {
			if n > curbest {
				curbest = n
			}
		} else {
			maxjoltageintl(n, base/10, v[i+1:])
		}
	}
}

func pow10(n int) int {
	x := 1
	for range n - 1 {
		x *= 10
	}
	return x
}
