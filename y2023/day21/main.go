package day21

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

type VisitMap map[int]int

func findStart(gr grid.Grid) grid.Position {
	return gr.FindPosition('S')
}

func expand(gr grid.Grid, positions VisitMap) VisitMap {
	vm := VisitMap{}
	for encodedPos, cnt := range positions {
		pos := gr.DecodePosition(encodedPos)
		for _, dir := range grid.FourSides {
			newPos := pos.Add(dir)
			if gr.Contains(newPos) && gr[newPos.Row][newPos.Col] != '#' {
				vm[gr.EncodePosition(newPos)] += cnt
			}
		}
	}
	return vm
}

func countVisited(gr grid.Grid, start grid.Position, steps int) int {
	ps := VisitMap{gr.EncodePosition(start): 1}
	for i := 0; i < steps; i++ {
		ps = expand(gr, ps)
	}
	return len(ps)
}

func SolveV1(input string) int {
	gr := utils.NonEmptyLines(input)
	start := findStart(gr)
	return countVisited(gr, start, 64)
}

func SolveV2(input string) int {
	gr := utils.NonEmptyLines(input)
	start := findStart(gr)

	steps := 26501365 // == 131 * 202300 + 65
	size := 131
	additionalSteps := steps % size

	return countVisited(gr, start, additionalSteps)
}
