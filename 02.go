package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

func main() {
	lines := Input(os.Args[1], "\n", true)
	part1 := 0
	part2 := 0
	for _, intvl := range Spac(lines[0], ",", -1) {
		v := Vatoi(Spac(intvl, "-", -1))
		for n := v[0]; n <= v[1]; n++ {
			if !isvalid(fmt.Sprintf("%d", n)) {
				part1 += n
			}
			if !isvalid2(fmt.Sprintf("%d", n)) {
				part2 += n
			}
		}
	}
	Sol(part1)
	Sol(part2)
}

func isvalid(x string) bool {
	return x[:len(x)/2] != x[len(x)/2:]
}

func isvalid2(x string) bool {
	for sz := 1; sz <= len(x)/2; sz++ {
		if len(x)%sz != 0 {
			continue
		}
		if !isvalid2intl(x, sz) {
			return false
		}
	}
	return true
}

func isvalid2intl(x string, sz int) bool {
	first := x[:sz]
	rest := x[sz:]
	for len(rest) >= sz {
		if rest[:sz] != first {
			return true
		}
		rest = rest[sz:]
	}
	return rest != ""
}
