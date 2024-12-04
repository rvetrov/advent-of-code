package day04

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

const xmas = "XMAS"
const mas = "MAS"

func contains(gr grid.Grid, pos grid.Position, dir grid.Direction, s string) bool {
	for i := 0; i < len(s); i++ {
		if next, ok := gr.At(pos); !ok || next != s[i] {
			return false
		}
		pos = pos.Add(dir)
	}
	return true
}

func SolveV1(input string) int {
	gr := grid.Grid(utils.NonEmptyLines(input))

	var res int
	for i := 0; i < gr.Rows(); i += 1 {
		for j := 0; j < gr.Cols(); j++ {
			for _, dir := range grid.EightSides {
				if contains(gr, grid.Position{Row: i, Col: j}, dir, xmas) {
					res += 1
				}
			}
		}
	}
	return res
}

func SolveV2(input string) int {
	gr := grid.Grid(utils.NonEmptyLines(input))

	crosses := []struct {
		firstDir       grid.Direction
		secondPosDelta grid.Direction
		secondDir      grid.Direction
	}{
		{grid.DownRight, grid.Right.Multiplied(2), grid.DownLeft},
		{grid.DownLeft, grid.Down.Multiplied(2), grid.UpLeft},
		{grid.UpLeft, grid.Left.Multiplied(2), grid.UpRight},
		{grid.UpRight, grid.Up.Multiplied(2), grid.DownRight},
	}

	var res int
	for i := 0; i < gr.Rows(); i += 1 {
		for j := 0; j < gr.Cols(); j++ {
			firstPos := grid.Position{Row: i, Col: j}
			for _, cross := range crosses {
				secondPos := firstPos.Add(cross.secondPosDelta)
				if contains(gr, firstPos, cross.firstDir, mas) && contains(gr, secondPos, cross.secondDir, mas) {
					res += 1
				}
			}
		}
	}
	return res
}
