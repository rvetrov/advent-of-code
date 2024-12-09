package executor

import (
	"adventofcode.com/y2024/day01"
	"adventofcode.com/y2024/day02"
	"adventofcode.com/y2024/day03"
	"adventofcode.com/y2024/day04"
	"adventofcode.com/y2024/day05"
	"adventofcode.com/y2024/day06"
	"adventofcode.com/y2024/day07"
	"adventofcode.com/y2024/day08"
	"adventofcode.com/y2024/day09"
)

var Y2024 = &Executor{
	Name: "2024",
	path: "y2024",
	solvers: map[string]Task{
		"day01": {[]taskSolver{day01.SolveV1, day01.SolveV2}},
		"day02": {[]taskSolver{day02.SolveV1, day02.SolveV2}},
		"day03": {[]taskSolver{day03.SolveV1, day03.SolveV2}},
		"day04": {[]taskSolver{day04.SolveV1, day04.SolveV2}},
		"day05": {[]taskSolver{day05.SolveV1, day05.SolveV2}},
		"day06": {[]taskSolver{day06.SolveV1, day06.SolveV2}},
		"day07": {[]taskSolver{day07.SolveV1, day07.SolveV2}},
		"day08": {[]taskSolver{day08.SolveV1, day08.SolveV2}},
		"day09": {[]taskSolver{day09.SolveV1, day09.SolveV2}},
	},
}
