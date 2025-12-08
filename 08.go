package main

import (
	. "aoc/util"
	"cmp"
	"math"
	"os"
	"slices"
)

type point struct {
	x, y, z float64
}

var points []point
var conn = map[[2]int]bool{}

type adist struct {
	d    float64
	i, j int
}

func main() {
	lines := Input(os.Args[1], "\n", true)

	for _, line := range lines {
		v := Vatoi(Spac(line, ",", -1))
		points = append(points, point{float64(v[0]), float64(v[1]), float64(v[2])})
	}

	distv := []adist{}

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			d := math.Sqrt(math.Pow(points[i].x-points[j].x, 2) + math.Pow(points[i].y-points[j].y, 2) + math.Pow(points[i].z-points[j].z, 2))
			distv = append(distv, adist{d, i, j})
		}
	}

	slices.SortFunc(distv, func(a, b adist) int { return cmp.Compare(a.d, b.d) })

	N := 10
	if len(lines) > 500 {
		N = 1000
	}

	comp0 := make(Set[int])
	comp0.Add(0)

	for cnt, ad := range distv {
		if cnt == N {
			seen := make(Set[int])
			countsv := []int{}
			for len(seen) < len(points) {
				for k := range points {
					if !seen[k] {
						sz := componentsz(k, seen)
						if sz != 1 {
							countsv = append(countsv, sz)
						}
					}
				}
			}
			slices.Sort(countsv)
			Pln(countsv)
			Sol(countsv[len(countsv)-1] * countsv[len(countsv)-2] * countsv[len(countsv)-3])
		}

		i, j := ad.i, ad.j
		//Pln("Joining", points[i], points[j], cnt)
		if i == 0 && j == 0 {
			break
		}

		conn[[2]int{i, j}] = true
		conn[[2]int{j, i}] = true

		switch {
		case comp0[i] && !comp0[j]:
			componentsz(j, comp0)
		case !comp0[i] && comp0[j]:
			componentsz(i, comp0)
		}

		if len(comp0) == len(points) {
			Pln("Final connection is", points[i], points[j])
			Sol(int(points[i].x * points[j].x))
			break
		}
	}
}

func componentsz(k int, seen Set[int]) int {
	comp := make(Set[int])
	q := make(Set[int])
	q.Add(k)
	for len(q) > 0 {
		cur := OneKey(q)
		delete(q, cur)
		if seen[cur] {
			continue
		}
		seen.Add(cur)
		comp.Add(cur)
		for j := range points {
			if conn[[2]int{cur, j}] && !seen[j] {
				q.Add(j)
			}
		}
	}
	return len(comp)
}
