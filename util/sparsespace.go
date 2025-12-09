package util

import (
	"fmt"
	"math"
	"sort"
)

type SparseSpace[T any] struct {
	coords [][]int // coords[0] sorted list of all x points, coords[1] sorted list of all y points, etc
	root   *ssNode[T]
}

type SSCell struct {
	Coord []int
	Sz    []int
}

type ssNode[T any] struct {
	children []*ssNode[T]
	leaves   []T
}

// NewSparseSpace creates a sparse space where only the coordinates are
// relevant. Each coordinate is a list of integer but all the element of
// coords must have the same length.
// Example: NewSparseSpace([][]int{ { x0, y0 }, { x1, y1} ... })
func NewSparseSpace[T any](coords [][]int) *SparseSpace[T] {
	ss := &SparseSpace[T]{}
	ss.coords = make([][]int, len(coords[0]))
	for _, coord := range coords {
		for i := range coord {
			ss.coords[i] = append(ss.coords[i], coord[i])
		}
	}
	for i := range ss.coords {
		ss.coords[i] = append(ss.coords[i], math.MinInt, math.MaxInt)
		sort.Ints(ss.coords[i])
		ss.coords[i] = uniq(ss.coords[i])
	}
	ss.root = newSSNode[T](ss.coords)
	return ss
}

func newSSNode[T any](coords [][]int) *ssNode[T] {
	r := &ssNode[T]{}
	if len(coords) == 1 {
		r.leaves = make([]T, len(coords[0]))
		return r
	}
	for range coords[0] {
		r.children = append(r.children, newSSNode[T](coords[1:]))
	}
	return r
}

// Suspace(start, end) iterates through all the cells between start and end (excluding end itself).
// Both start and end must have been specified as valid coordinates when creating the SparseSpace.
// Exmaple:
//
//	for cell, p := range ss.Subspace(start, end) {
//		p is a pointer to the cell value
//		cell describes the cell itself
//	}
func (ss *SparseSpace[T]) Subspace(start, end []int) func(func(SSCell, *T) bool) {
	return func(yield func(SSCell, *T) bool) {
		sidxs, eidxs := make([]int, 0), make([]int, 0)
		for i := range start {
			sidx, eidx := ss.toindex(start[i], end[i], i)
			sidxs = append(sidxs, sidx)
			eidxs = append(eidxs, eidx)
		}
		ss.root.subspace(sidxs, eidxs, ss, 0, make([]int, len(ss.coords)), make([]int, len(ss.coords)), yield)
	}
}

func (ss *SparseSpace[T]) toindex(start, end int, i int) (int, int) {
	var s int
	found := false
	for s = range ss.coords[i] {
		if ss.coords[i][s] == start {
			found = true
			break
		}
		if ss.coords[i][s] > start {
			break
		}
	}
	if !found {
		panic(fmt.Errorf("coord not found %d", start))
	}
	found = false
	var e int
	for e = s + 1; e < len(ss.coords[i]); e++ {
		if ss.coords[i][e] == end {
			found = true
			break
		}
		if ss.coords[i][e] > end {
			break
		}
	}
	if !found {
		panic(fmt.Errorf("coord not found %d", end))
	}
	return s, e
}

func (ssn *ssNode[T]) subspace(start, end []int, ss *SparseSpace[T], depth int, coord, sz []int, yield func(SSCell, *T) bool) bool {
	if ssn.children == nil {
		for i := start[0]; i < end[0]; i++ {
			coord[depth] = ss.coords[depth][i]
			sz[depth] = ss.coords[depth][i+1] - ss.coords[depth][i]
			if !yield(SSCell{Coord: coord, Sz: sz}, &ssn.leaves[i]) {
				return false
			}
		}
		return true
	}

	for i := start[0]; i < end[0]; i++ {
		coord[depth] = ss.coords[depth][i]
		sz[depth] = ss.coords[depth][i+1] - ss.coords[depth][i]
		if !ssn.children[i].subspace(start[1:], end[1:], ss, depth+1, coord, sz, yield) {
			return false
		}
	}
	return true
}

func uniq(v []int) []int {
	r := make([]int, 0, len(v))
	for _, x := range v {
		if len(r) == 0 || r[len(r)-1] != x {
			r = append(r, x)
		}
	}
	return r
}

func (ss *SparseSpace[T]) Containing(p []int) (SSCell, *T) {
	ssn := ss.root
	coord := []int{}
	sz := []int{}
	for i := range p {
		found := false
		for j := range ss.coords[i] {
			if p[i] < ss.coords[i][j] {
				coord = append(coord, ss.coords[i][j-1])
				sz = append(sz, ss.coords[i][j]-ss.coords[i][j-1])
				if ssn.children == nil {
					return SSCell{coord, sz}, &ssn.leaves[j-1]
				} else {
					found = true
					ssn = ssn.children[j-1]
					break
				}
			}
		}
		if !found {
			panic("not found")
		}
	}
	panic("not found")
}
