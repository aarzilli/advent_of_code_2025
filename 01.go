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
		Pln("moving", instr)
		sign := 0
		switch instr[0] {
		case 'L':
			sign = -1
		case 'R':
			sign = +1
		default:
			panic("blah")
		}
		n := sign*Atoi(instr[1:])
		prevpos := pos
		pos += n
		pos %= 100
		if pos < 0 {
			pos += 100
		}
		Pln(pos)
		if pos == 0 {
			part1++
		}
		sign2 := 0
		if pos - prevpos < 0 {
			sign2 = -1
		} else {
			sign2 = +1
		}
		if sign2 != sign {
			Pln("\tcrossed")
		}
		p2 := prevpos
		for range Atoi(instr[1:]) {
			p2 += sign
			if p2 < 0 {
				p2 = 99
			}
			if p2 == 100 {
				p2 = 0
			}
			if p2 == 0 {
				Pln("\tcrossed2")
				part2++
			}
		}
		if p2 != pos {
			panic("mismatch")
		}
	}
	Sol(part1)
	Sol(part2)
}
