package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

var adj = map[string]map[string]bool{}

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	for _, line := range lines {
		v := Spac(line, " ", -1)
		v[0] = v[0][:len(v[0])-1]
		adj[v[0]] = make(map[string]bool)
		for _, n := range v[1:] {
			adj[v[0]][n] = true
		}
	}

	fh, err := os.Create("11.dot")
	Must(err)
	fmt.Fprintf(fh, "digraph blah {\n")
	for k := range adj {
		for nb := range adj[k] {
			fmt.Fprintf(fh, "%s -> %s\n", k, nb)
		}
	}
	fmt.Fprintf(fh, "}\n")

	to := toposort()
	Sol(countpaths(to, "you")["out"])

	var dacpos, fftpos int
	for i, k := range to {
		if k == "dac" {
			dacpos = i
		}
		if k == "fft" {
			fftpos = i
		}
	}

	if fftpos >= dacpos {
		panic("not implemented")
	}

	fromsvr := countpaths(to, "svr")
	fromfft := countpaths(to, "fft")
	fromdac := countpaths(to, "dac")

	Pln(fromsvr["fft"], fromfft["dac"], fromdac["out"])
	Sol(fromsvr["fft"] * fromfft["dac"] * fromdac["out"])
}

func toposort() []string {
	r := []string{}

	inc := make(map[string]map[string]bool)
	for a := range adj {
		for b := range adj[a] {
			if inc[b] == nil {
				inc[b] = make(map[string]bool)
			}
			inc[b][a] = true
		}
	}

	S := make(Set[string])

	for n := range adj {
		if len(inc[n]) == 0 {
			S.Add(n)
		}
	}

	for len(S) > 0 {
		n := OneKey(S)
		delete(S, n)
		r = append(r, n)
		for m := range adj[n] {
			delete(inc[m], n)
			if len(inc[m]) == 0 {
				S.Add(m)
			}
		}
	}

	return r
}

func countpaths(topoorder []string, start string) map[string]int {
	cnt := make(map[string]int)
	cnt[start] = 1
	for _, k := range topoorder {
		for nb := range adj[k] {
			cnt[nb] += cnt[k]
		}
	}
	return cnt
}
