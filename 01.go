package main

import (
	. "aoc/util"
	"os"
)

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	pos := 50
	part1 := 0
	part2 := 0
	for _, instr := range lines {
		sign := 0
		switch instr[0] {
		case 'L':
			sign = -1
		case 'R':
			sign = +1
		default:
			panic("blah")
		}
		n := sign * Atoi(instr[1:])
		prevpos := pos
		pos += n
		pos %= 100
		if pos < 0 {
			pos += 100
		}
		if pos == 0 {
			part1++
		}
		sign2 := 0
		if pos-prevpos < 0 {
			sign2 = -1
		} else {
			sign2 = +1
		}
		if sign2 != sign && prevpos != 0 {
			part2++
		} else if pos == 0 {
			part2++
		}
		part2 += Abs(n) / 100
	}
	Expect(992)
	Sol(part1)
	Expect(6133)
	Sol(part2)
}
