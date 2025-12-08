package main

import (
	. "aoc/util"
	"os"
	"math"
	_ "sort"
)

type point struct {
	x, y, z float64
}

var points []point
var dist = map[int]map[int]float64{}
var color []int
var nextcolor = 1

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	
	for _, line := range lines {
		v := Vatoi(Spac(line, ",", -1))
		points = append(points, point{ float64(v[0]), float64(v[1]), float64(v[2]) })
	}
	
	for i := range points {
		for j := i+1; j < len(points); j++ {
			d := math.Sqrt(math.Pow(points[i].x - points[j].x, 2) + math.Pow(points[i].y - points[j].y, 2) + math.Pow(points[i].z - points[j].z, 2))
			if dist[i] == nil {
				dist[i] = make(map[int]float64)
			}
			if dist[j] == nil {
				dist[j] = make(map[int]float64)
			}
			dist[i][j] = d
			dist[j][i] = d
		}
	}
	
	Pln("distances done")
	
	color = make([]int, len(points))
	
	prevmin := 0.0
	
	/*N := 10
	if len(lines) > 500 {
		Pln("big input")
		N = 1000
	}*/
	
	for cnt := 0; true; cnt++ {
		i, j := findmin(prevmin)
		Pln("Joining", points[i], points[j], cnt)
		prevmin = dist[i][j]
		if i == 0 && j == 0 {
			break
		}
		c := 0
		if color[i] != 0 && color[j] != 0 {
			c = color[i]
			c2 := color[j]
			for k := range color {
				if color[k] == c2 {
					color[k] = c
				}
			}
		} else if color[i] != 0 {
			c = color[i]
		} else if color[j] != 0 {
			c = color[j]
		} else {
			c = nextcolor
			nextcolor++
		}
		color[i] = c
		color[j] = c
		
		allsame := true
		c2 := color[0]
		for i := range color {
			if color[i] != c2 {
				allsame = false
				break
			}
		}
		if allsame {
			Pln("first connection is", points[i], points[j])
			Sol(points[i].x * points[j].x)
			break
		}
	}
	/*
	Pln(color)
	counts := map[int]int{}
	for i := range color {
		counts[color[i]]++
	}
	countsv := []int{}
	for k, v := range counts {
		if k == 0 {
			continue
		}
		countsv = append(countsv, v)
	}
	Pln(counts)
	sort.Ints(countsv)
	Pln(countsv)
	Sol(countsv[len(countsv)-1] * countsv[len(countsv)-2] * countsv[len(countsv)-3])*/
}

func findmin(prevmin float64) (int, int) {
	minpair := [2]int{ 0, 0 }
	first := true
	for i := range points {
		for j := i+1; j < len(points); j++ {
			if dist[i][j] > prevmin && (dist[i][j] < dist[minpair[0]][minpair[1]] || first) {
				first = false
				minpair = [2]int{ i, j }
			}
		}
	}
	return minpair[0], minpair[1]
}