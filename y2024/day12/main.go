package day12

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

func walk(pos grid.Position, gr grid.Grid, visited map[grid.Position]struct{}) (area, perimeter, sides int) {
	area++
	visited[pos] = struct{}{}
	curCh, _ := gr.At(pos)
	for _, dir := range grid.FourSides {
		nextPos := pos.Add(dir)
		if ch, inside := gr.At(nextPos); inside && ch == curCh {
			if _, v := visited[nextPos]; !v {
				nextArea, nextPerimeter, newSides := walk(nextPos, gr, visited)
				area += nextArea
				perimeter += nextPerimeter
				sides += newSides
			}
		} else {
			perimeter++
			nextToCur := pos.Add(dir.TurnCW())
			nextToNext := nextPos.Add(dir.TurnCW())
			nextToCurCh, nextToCurInside := gr.At(nextToCur)
			nextToNextCh, nextToNextInside := gr.At(nextToNext)
			if !(nextToCurInside && nextToCurCh == curCh) || (nextToNextInside && nextToNextCh == curCh) {
				sides++
			}
		}
	}
	return area, perimeter, sides
}

func SolveV1(input string) int {
	gr := grid.New(utils.NonEmptyLines(input))
	visited := make(map[grid.Position]struct{})
	res := 0
	for pos := gr.First(); gr.Contains(pos); pos = gr.Next(pos) {
		if _, v := visited[pos]; !v {
			area, perimeter, _ := walk(pos, gr, visited)
			res += area * perimeter
		}
	}
	return res
}

func SolveV2(input string) int {
	gr := grid.New(utils.NonEmptyLines(input))
	visited := make(map[grid.Position]struct{})
	res := 0
	for pos := gr.First(); gr.Contains(pos); pos = gr.Next(pos) {
		if _, v := visited[pos]; !v {
			area, _, sides := walk(pos, gr, visited)
			res += area * sides
		}
	}
	return res
}
