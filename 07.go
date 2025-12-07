package main

import (
	. "aoc/util"
	"os"
)

var M [][]byte

type position struct {
	i, j int
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}
	var s position
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 'S' {
				s = position{i, j}
				break
			}
		}
	}

	ps := map[position]int{}
	ps[s] = 1

	part1 := 0
	part2 := 0
	for len(ps) > 0 {
		psnext := map[position]int{}
		for p, cnt := range ps {
			p.i++
			if p.i >= len(M) {
				part2 += cnt
				continue
			}
			if M[p.i][p.j] == '^' {
				part1++
				psnext[position{p.i, p.j - 1}] += cnt
				psnext[position{p.i, p.j + 1}] += cnt
			} else {
				psnext[p] += cnt
			}
		}
		ps = psnext
	}
	Sol(part1)
	Sol(part2)
}
