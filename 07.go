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
				s = position{ i, j }
				break
			}
		}
	}
	
	ps := make(Set[position])
	ps.Add(s)
	
	part1 := 0
	part2 := 1
	for len(ps) > 0 {
		psnext := make(Set[position])
		for p := range ps {
			p.i++
			if p.i >= len(M) {
				continue
			}
			if M[p.i][p.j] == '^' {
				part1++
				psnext.Add(position{ p.i, p.j-1 })
				psnext.Add(position{ p.i, p.j+1 })
			} else {
				psnext.Add(p)
			}
		}
		ps = psnext
	}
	Sol(part1)
	Sol(part2)
}
