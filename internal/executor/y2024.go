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
	"adventofcode.com/y2024/day10"
	"adventofcode.com/y2024/day11"
	"adventofcode.com/y2024/day12"
	"adventofcode.com/y2024/day13"
	"adventofcode.com/y2024/day14"
	"adventofcode.com/y2024/day15"
	"adventofcode.com/y2024/day16"
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
		"day10": {[]taskSolver{day10.SolveV1, day10.SolveV2}},
		"day11": {[]taskSolver{day11.SolveV1, day11.SolveV2}},
		"day12": {[]taskSolver{day12.SolveV1, day12.SolveV2}},
		"day13": {[]taskSolver{day13.SolveV1, day13.SolveV2}},
		"day14": {[]taskSolver{day14.SolveV1, day14.SolveV2}},
		"day15": {[]taskSolver{day15.SolveV1, day15.SolveV2}},
		"day16": {[]taskSolver{day16.SolveV1, day16.SolveV2}},
	},
}
