package day06

import (
	"slices"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

func walk(
	gr grid.Grid, pos grid.Position, dir grid.Direction, visited map[grid.Position][]grid.Direction,
) (steps int, loopFound bool) {

	madeStep := 0
	if visitedDirs, isVisited := visited[pos]; isVisited {
		if slices.Contains(visitedDirs, dir) {
			return 0, true
		}
	} else {
		madeStep += 1
	}
	visited[pos] = append(visited[pos], dir)

	next := pos.Add(dir)
	if nextCh, ok := gr.At(next); !ok {
		return madeStep, false
	} else if nextCh == '#' {
		dir = dir.TurnCW()
	} else {
		pos = next
	}

	steps, loopFound = walk(gr, pos, dir, visited)
	return steps + madeStep, loopFound
}

func SolveV1(input string) int {
	gr := grid.New(utils.NonEmptyLines(input))
	start, _ := gr.FindPosition('^')
	dir := grid.Up
	visited := map[grid.Position][]grid.Direction{}

	steps, _ := walk(gr, start, dir, visited)
	return steps
}

func SolveV2(input string) int {
	gr := grid.New(utils.NonEmptyLines(input))
	start, _ := gr.FindPosition('^')
	dir := grid.Up

	defaultVisited := map[grid.Position][]grid.Direction{}
	walk(gr, start, dir, defaultVisited)

	res := 0
	for pos := gr.First(); gr.Contains(pos); pos = gr.Next(pos) {
		if _, onDefaultPath := defaultVisited[pos]; !onDefaultPath {
			continue
		}
		if ch, _ := gr.At(pos); ch != '.' {
			continue
		}
		gr.SetAt(pos, '#')
		visited := map[grid.Position][]grid.Direction{}
		if _, loopFound := walk(gr, start, dir, visited); loopFound {
			res += 1
		}
		gr.SetAt(pos, '.')
	}

	return res
}
