package executor

import (
	"adventofcode.com/y2024/day01"
)

var Y2024 = &Executor{
	Name: "2024",
	path: "y2024",
	solvers: map[string]Task{
		"day01": {[]taskSolver{day01.SolveV1, day01.SolveV2}},
	},
}
