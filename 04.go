package main

import (
	. "aoc/util"
	"os"
)

var M [][]byte

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}

	part1 := 0
	for i := range M {
		for j := range M[i] {
			if get(i, j) == 1 {
				tot := get(i-1, j-1) + get(i-1, j) + get(i-1, j+1) +
					get(i, j-1) + get(i, j+1) +
					get(i+1, j-1) + get(i+1, j) + get(i+1, j+1)
				if tot < 4 {
					part1++
				}
			}
		}
	}
	Sol(part1)

	part2 := 0
	for {
		removed := step()
		part2 += removed
		if removed == 0 {
			break
		}

		/*
			for i := range M {
				Pln(string(M[i]))
			}
			Pln()
			Pln()
		*/
	}

	Sol(part2)
}

func get(i, j int) int {
	if i < 0 || i >= len(M) {
		return 0
	}
	if j < 0 || j >= len(M[i]) {
		return 0
	}
	if M[i][j] == '@' {
		return 1
	}
	return 0
}

func step() int {
	toremove := [][2]int{}
	for i := range M {
		for j := range M[i] {
			if get(i, j) == 1 {
				tot := get(i-1, j-1) + get(i-1, j) + get(i-1, j+1) +
					get(i, j-1) + get(i, j+1) +
					get(i+1, j-1) + get(i+1, j) + get(i+1, j+1)
				if tot < 4 {
					toremove = append(toremove, [2]int{i, j})
				}
			}
		}
	}

	for _, p := range toremove {
		M[p[0]][p[1]] = '.'
	}
	return len(toremove)
}
