package executor

import (
	"golang.org/x/exp/maps"

	"adventofcode.com/y2023/day01"
	"adventofcode.com/y2023/day02"
	"adventofcode.com/y2023/day04"
	"adventofcode.com/y2023/day05"
	"adventofcode.com/y2023/day06"
	"adventofcode.com/y2023/day07"
	"adventofcode.com/y2023/day08"
	"adventofcode.com/y2023/day09"
	"adventofcode.com/y2023/day10"
	"adventofcode.com/y2023/day11"
	"adventofcode.com/y2023/day12"
	"adventofcode.com/y2023/day13"
	"adventofcode.com/y2023/day14"
	"adventofcode.com/y2023/day15"
	"adventofcode.com/y2023/day16"
	"adventofcode.com/y2023/day17"
	"adventofcode.com/y2023/day18"
)

type taskSolver func(string) int

type Task struct {
	solvers []taskSolver
}

var tasks2023 = map[string]Task{
	"day01": {[]taskSolver{day01.SolveV1, day01.SolveV2}},
	"day02": {[]taskSolver{day02.SolveV1, day02.SolveV2}},
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
	"day17": {[]taskSolver{day17.SolveV1, day17.SolveV2}},
	"day18": {[]taskSolver{day18.SolveV1, day18.SolveV2}},
}

var KnownTasks2023 = maps.Keys(tasks2023)
