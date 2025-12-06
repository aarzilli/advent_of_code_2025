package main

import (
	. "aoc/util"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	M := make([][]string, len(lines))
	for i := range lines {
		M[i] = Noempty(Spac(lines[i], " ", -1))
		Pf("%q\n", M[i])
	}

	part1 := 0
	for j := range len(M[0]) {
		acc := 0
		if M[len(M)-1][j] == "*" {
			acc = 1
		}
		for i := 0; i < len(M)-1; i++ {
			switch M[len(M)-1][j] {
			case "*":
				acc *= Atoi(M[i][j])
			case "+":
				acc += Atoi(M[i][j])
			default:
				panic("blah")
			}
		}
		part1 += acc
	}
	Sol(part1)

	buf, err := ioutil.ReadFile(os.Args[1])
	Must(err)
	lines = strings.SplitN(string(buf), "\n", -1)

	part2 := 0
	nums := []int{}
	for j := len(lines[0]) - 1; j >= 0; j-- {
		acc := 0
		for i := 0; i < len(lines); i++ {
			if len(lines[i]) == 0 {
				continue
			}
			//Pf("accessing %q at %d %d %c\n", lines[i], i, j, lines[i][j])
			switch lines[i][j] {
			case '*':
				nums = append(nums, acc)
				acc = 0
				tot := 1
				for i := range nums {
					tot *= nums[i]
				}
				Pln("total:", tot)
				nums = nums[:0]
				part2 += tot
			case '+':
				nums = append(nums, acc)
				acc = 0
				tot := 0
				for i := range nums {
					tot += nums[i]
				}
				Pln("total", tot)
				nums = nums[:0]
				part2 += tot
			default:
				if lines[i][j] >= '0' && lines[i][j] <= '9' {
					acc *= 10
					acc += int(lines[i][j] - '0')
				}
			}
		}
		if acc != 0 {
			nums = append(nums, acc)
		}
	}
	Sol(part2)
}
