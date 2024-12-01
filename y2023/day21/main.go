package day21

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

type VisitMap map[int]int

func findStart(gr grid.Grid) grid.Position {
	n, m := gr.Rows(), gr.Cols()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if gr[i][j] == 'S' {
				return grid.Position{Row: i, Col: j}
			}
		}
	}
	return grid.Position{}
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
	/*
		leftUpperCorner := grid.Position{Row: 0, Col: 0}
		leftLowerCorner := grid.Position{Row: gr.Rows() - 1, Col: 0}
		rightUpperCorner := grid.Position{Row: 0, Col: gr.Cols() - 1}
		rightLowerCorner := grid.Position{Row: gr.Rows() - 1, Col: gr.Cols() - 1}

		steps := 26501365 // 26501365 == 131 * 202300 + 65
		size := 131

		x := (steps - 66) % 131
		additionalSteps := steps % size
	*/
	gr := utils.NonEmptyLines(input)
	start := findStart(gr)

	steps := 26501365 // 26501365 == 131 * 202300 + 65
	size := 131
	additionalSteps := steps % size

	return countVisited(gr, start, additionalSteps)
}
