package executor

import (
	"adventofcode.com/y2024/day01"
	"adventofcode.com/y2024/day02"
	"adventofcode.com/y2024/day03"
	"adventofcode.com/y2024/day04"
)

var Y2024 = &Executor{
	Name: "2024",
	path: "y2024",
	solvers: map[string]Task{
		"day01": {[]taskSolver{day01.SolveV1, day01.SolveV2}},
		"day02": {[]taskSolver{day02.SolveV1, day02.SolveV2}},
		"day03": {[]taskSolver{day03.SolveV1, day03.SolveV2}},
		"day04": {[]taskSolver{day04.SolveV1, day04.SolveV2}},
	},
}
