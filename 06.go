package main

import (
	. "aoc/util"
	"os"
)

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	M := make([][]string, len(lines))
	for i := range lines {
		M[i] = Noempty(Spac(lines[i], " ", -1))
		Pf("%q\n", M[i])
	}

	part1 := 0
	for j := range len(M[0]) {
		acc := 0
		if M[len(M)-1][j] == "*" {
			acc = 1
		}
		for i := 0; i < len(M)-1; i++ {
			switch M[len(M)-1][j] {
			case "*":
				acc *= Atoi(M[i][j])
			case "+":
				acc += Atoi(M[i][j])
			default:
				panic("blah")
			}
		}
		part1 += acc
	}
	Sol(part1)
}
