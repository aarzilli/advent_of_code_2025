package main

import (
	. "aoc/util"
	"os"
	"slices"
)

type rng struct {
	start, end int
}

var ranges []rng

func main() {
	lines := Input(os.Args[1], "\n\n", true)

	for _, s := range Spac(lines[0], "\n", -1) {
		v := Vatoi(Spac(s, "-", -1))
		ranges = append(ranges, rng{v[0], v[1]})
	}

	cases := Vatoi(Spac(lines[1], "\n", -1))

	part1 := 0
	for _, s := range cases {
		fresh := false
		for _, rng := range ranges {
			if rng.contains(s) {
				fresh = true
				break
			}
		}
		if fresh {
			part1++
		}
	}
	Sol(part1)

	slices.SortFunc(ranges, func(a, b rng) int {
		return a.start - b.start
	})

	for {
		changed := false
		for i := range ranges {
			if ranges[i].start == -1 {
				continue
			}
			for j := i + 1; j < len(ranges); j++ {
				if ranges[j].start == -1 {
					continue
				}
				if ranges[i].contains(ranges[j].start) {
					if ranges[j].end > ranges[i].end {
						ranges[i].end = ranges[j].end
					}
					ranges[j].start = -1
					ranges[j].end = -1
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}
	Pln(ranges)
	part2 := 0
	for _, rng := range ranges {
		if rng.start == -1 {
			continue
		}
		part2 += rng.end - rng.start + 1
	}
	Sol(part2)
}

func (rng *rng) contains(x int) bool {
	return x >= rng.start && x <= rng.end
}

// 338186538605731 wrong