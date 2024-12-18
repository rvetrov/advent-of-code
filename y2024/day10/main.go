package day10

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

func trailheadsV1(gr grid.Grid, pos grid.Position, mark int, visited [][]int) int {
	visited[pos.Row][pos.Col] = mark
	cur, _ := gr.At(pos)
	if cur == '9' {
		return 1
	}
	res := 0
	for _, dir := range grid.FourSides {
		newPos := pos.Add(dir)
		if newCh, inside := gr.At(newPos); inside && newCh == cur+1 && visited[newPos.Row][newPos.Col] != mark {
			res += trailheadsV1(gr, newPos, mark, visited)
		}
	}
	return res
}

func trailheadsV2(gr grid.Grid, pos grid.Position) int {
	cur, _ := gr.At(pos)
	if cur == '9' {
		return 1
	}
	res := 0
	for _, dir := range grid.FourSides {
		newPos := pos.Add(dir)
		if newCh, inside := gr.At(newPos); inside && newCh == cur+1 {
			res += trailheadsV2(gr, newPos)
		}
	}
	return res
}

func SolveV1(input string) int {
	gr := grid.New(utils.NonEmptyLines(input))
	var visited [][]int
	for range gr.Rows() {
		visited = append(visited, make([]int, gr.Cols()))
	}

	res := 0
	for pos := gr.First(); gr.Contains(pos); pos = gr.Next(pos) {
		if ch, _ := gr.At(pos); ch == '0' {
			mark := gr.EncodePosition(pos) + 1
			res += trailheadsV1(gr, pos, mark, visited)
		}
	}

	return res
}

func SolveV2(input string) int {
	gr := grid.New(utils.NonEmptyLines(input))
	res := 0
	for pos := gr.First(); gr.Contains(pos); pos = gr.Next(pos) {
		if ch, _ := gr.At(pos); ch == '0' {
			res += trailheadsV2(gr, pos)
		}
	}
	return res
}
