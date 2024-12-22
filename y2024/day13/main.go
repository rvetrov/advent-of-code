package day13

import (
	"fmt"
	"strings"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/math"
	"adventofcode.com/internal/utils"
)

func parseVec(line, formatStr string) grid.Direction {
	_, vecStr, _ := strings.Cut(line, ": ")
	x, y := 0, 0
	if read, err := fmt.Sscanf(vecStr, formatStr, &x, &y); read != 2 || err != nil {
		panic(vecStr)
	}
	return grid.Direction{DR: y, DC: x}
}

func parseInput(lines []string) (a, b, target grid.Direction) {
	a = parseVec(lines[0], "X+%d, Y+%d")
	b = parseVec(lines[1], "X+%d, Y+%d")
	target = parseVec(lines[2], "X=%d, Y=%d")
	return a, b, target
}

func bruteForceSolve(a, b, target grid.Direction) int {
	res := 0
	for ai := range 101 {
		for bi := range 101 {
			cost := ai*3 + bi
			if (res == 0 || res > cost) && a.Multiplied(ai).Add(b.Multiplied(bi)) == target {
				res = cost
			}
		}
	}
	return res
}

func solve(a, b, target grid.Direction) int {
	/*
		a.DC * m + b.DC * n = target.DC
		a.DR * m + b.DR * n = target.DR

		a1*m + a2*n = a3
		b1*m + b2*n = b3
	*/
	m, n := 0, 0

	a1, a2, a3 := a.DC, b.DC, target.DC
	b1, b2, b3 := a.DR, b.DR, target.DR
	lcm := math.LCM(a1, b1)

	k1 := lcm / a1
	a1 *= k1
	a2 *= k1
	a3 *= k1

	k2 := lcm / b1
	b1 *= k2
	b2 *= k2
	b3 *= k2

	b1 -= a1
	b2 -= a2
	b3 -= a3

	if b3%b2 != 0 {
		return 0
	}
	n = b3 / b2

	a3 -= a2 * n
	if a3%a1 != 0 {
		return 0
	}
	m = a3 / a1

	if m < 0 || n < 0 {
		return 0
	}
	return m*3 + n
}

func SolveV1(input string) int {
	res := 0
	for _, block := range utils.EmptyLineSeparatedBlocks(input) {
		a, b, target := parseInput(block)
		res += bruteForceSolve(a, b, target)
	}
	return res
}

func SolveV2(input string) int {
	res := 0
	for _, block := range utils.EmptyLineSeparatedBlocks(input) {
		a, b, target := parseInput(block)
		target.DC += 10000000000000
		target.DR += 10000000000000
		res += solve(a, b, target)
	}
	return res
}
