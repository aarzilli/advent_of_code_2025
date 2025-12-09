package main

import (
	. "aoc/util"
	"os"
)

type point struct {
	x, y int
}

var points []point
var ss *SparseSpace[bool]

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	for _, line := range lines {
		v := Vatoi(Spac(line, ",", -1))
		points = append(points, point{x: v[0], y: v[1]})
	}
	maxarea := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			a := (Abs(points[i].x-points[j].x) + 1) * (Abs(points[i].y-points[j].y) + 1)
			if a >= maxarea {
				maxarea = a
			}
		}
	}
	Sol(maxarea)

	coords := [][]int{}
	for i := range points {
		coords = append(coords, []int{points[i].x, points[i].y}, []int{points[i].x + 1, points[i].y + 1})
	}
	Pln(coords)
	ss = NewSparseSpace[bool](coords)

	for i := 0; i < len(points)-1; i++ {
		if points[i].x != points[i+1].x && points[i].y != points[i+1].y {
			panic("blah")
		}
		drawline(points[i], points[i+1])
	}
	drawline(points[len(points)-1], points[0])

	displayexample()

	miny := points[0].y
	for _, p := range points {
		if p.y < miny {
			miny = p.y
		}
	}

	q := make(Set[point])
	for _, p := range points {
		if goodstart(p) && p.y == miny {
			q.Add(point{x: p.x + 1, y: p.y + 1})
			break
		}
	}

	for len(q) > 0 {
		p := OneKey(q)
		delete(q, p)
		//Pln(p)
		cell, val := ss.Containing([]int{p.x, p.y})
		if *val == true {
			continue
		}
		*val = true
		q.Add(point{cell.Coord[0], cell.Coord[1] + cell.Sz[1]})
		q.Add(point{cell.Coord[0], cell.Coord[1] - 1})
		q.Add(point{cell.Coord[0] + cell.Sz[0], cell.Coord[1]})
		q.Add(point{cell.Coord[0] - 1, cell.Coord[1]})
		//displayexample()
	}

	maxarea2 := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]
			area := (Abs(a.x-b.x) + 1) * (Abs(a.y-b.y) + 1)
			if area < maxarea2 {
				continue
			}
			allinside := true
			for _, pval := range ss.Subspace([]int{min(a.x, b.x), min(a.y, b.y)}, []int{max(a.x, b.x) + 1, max(a.y, b.y) + 1}) {
				if *pval != true {
					allinside = false
					break
				}
			}
			if !allinside {
				continue
			}
			maxarea2 = area
		}
	}
	Sol(maxarea2)
}

func displayexample() {
	for y := range 9 {
		for x := range 15 {
			_, p := ss.Containing([]int{x, y})
			if *p {
				Pf("#")
			} else {
				Pf(".")
			}
		}
		Pln()
	}
	Pln()
}

func drawline(a, b point) {
	Pln(a, b)
	for _, p := range ss.Subspace([]int{min(a.x, b.x), min(a.y, b.y)},
		[]int{max(a.x, b.x) + 1, max(a.y, b.y) + 1}) {
		*p = true
	}
}

func goodstart(p point) bool {
	if _, p := ss.Containing([]int{p.x + 1, p.y + 1}); *p != false {
		return false
	}
	if _, p := ss.Containing([]int{p.x + 1, p.y}); *p != true {
		return false
	}
	if _, p := ss.Containing([]int{p.x, p.y + 1}); *p != true {
		return false
	}
	return true
}
