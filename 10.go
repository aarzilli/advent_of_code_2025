package main

import (
	. "aoc/util"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type machine struct {
	pat  string
	btns [][]int
	jtgt []int
}

var machines []machine

func main() {
	lines := Input(os.Args[1], "\n", true)

	for _, line := range lines {
		v := Spac(line, " ", -1)

		for i := range v {
			v[i] = v[i][1:]
			v[i] = v[i][:len(v[i])-1]
		}

		var m machine

		m.pat = v[0]

		m.jtgt = Vatoi(Spac(v[len(v)-1], ",", -1))
		v = v[:len(v)-1]

		v = v[1:]

		for i := range v {
			m.btns = append(m.btns, Vatoi(Spac(v[i], ",", -1)))
		}

		machines = append(machines, m)
	}

	part1 := 0
	for _, m := range machines {
		n := search(m)
		part1 += n
	}
	Sol(part1)

	fh, err := os.Create("10.wls")
	Must(err)
	fmt.Fprintf(fh, "#!/usr/bin/env wolframscript\n")
	for _, m := range machines {
		printsystem(fh, m)
	}
	Must(fh.Close())

	out, err := exec.Command("wolframscript", "10.wls").CombinedOutput()
	Must(err)

	v := Vatoi(Spac(strings.TrimSpace(string(out)), "\n", -1))
	Pln(v)
	Sol(Sum(v))
}

func search(m machine) int {
	djk := NewDijkstra[string](makelights(len(m.pat)))

	var cur string
	for djk.PopTo(&cur) {
		if cur == m.pat {
			return djk.Dist[cur]
		}

		for _, btn := range m.btns {
			djk.Add(cur, switchlights(cur, btn), 1)
		}
	}

	panic("not found")
}

func makelights(len int) string {
	v := make([]byte, len)
	for i := range v {
		v[i] = '.'
	}
	return string(v)
}

func switchlights(in string, btn []int) string {
	v := []byte(in)
	for _, i := range btn {
		if v[i] == '.' {
			v[i] = '#'
		} else {
			v[i] = '.'
		}
	}
	return string(v)
}

func printsystem(fh io.Writer, m machine) {
	fmt.Fprintf(fh, "Print@MinValue[{")
	for i := range m.btns {
		fmt.Fprintf(fh, "a%d", i)
		if i != len(m.btns)-1 {
			fmt.Fprintf(fh, "+")
		}
	}
	fmt.Fprintf(fh, ", ")
	for i := range m.jtgt {
		fmt.Fprintf(fh, "%d == ", m.jtgt[i])
		var l []string
		for k := range m.btns {
			found := false
			for j := range m.btns[k] {
				if m.btns[k][j] == i {
					found = true
					break
				}
			}
			if found {
				l = append(l, fmt.Sprintf("a%d", k))
			}
		}
		fmt.Fprintf(fh, "%s", strings.Join(l, " + "))
		fmt.Fprintf(fh, " && ")
	}
	for i := range m.btns {
		fmt.Fprintf(fh, "a%d >= 0", i)
		if i != len(m.btns)-1 {
			fmt.Fprintf(fh, " && ")
		}
	}
	fmt.Fprintf(fh, "}, {")
	for i := range m.btns {
		fmt.Fprintf(fh, "a%d", i)
		if i != len(m.btns)-1 {
			fmt.Fprintf(fh, ", ")
		}
	}
	fmt.Fprintf(fh, "}, Integers]\n")
}
