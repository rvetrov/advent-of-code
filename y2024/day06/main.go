package day06

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
	"slices"
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
	prevPos := pos
	prevVisited := visited[pos]
	visited[pos] = append(prevVisited, dir)

	next := pos.Add(dir)
	if nextCh, ok := gr.At(next); !ok {
		return madeStep, false
	} else if nextCh == '#' {
		dir = dir.TurnCW()
	} else {
		pos = next
	}

	steps, loopFound = walk(gr, pos, dir, visited)
	visited[prevPos] = prevVisited
	return steps + madeStep, loopFound
}

func walkAndFindLoops(gr grid.Grid, pos grid.Position, dir grid.Direction, visited map[grid.Position][]grid.Direction) int {
	if visitedDirs, isVisited := visited[pos]; isVisited {
		if slices.Contains(visitedDirs, dir) {
			return 1
		}
	}
	visited[pos] = append(visited[pos], dir)

	next := pos.Add(dir)
	if nextCh, ok := gr.At(next); !ok {
		return 0
	} else if nextCh == '#' {
		dir = dir.TurnCW()
	} else {
		// TODO
		pos = next
	}

	return walkAndFindLoops(gr, pos, dir, visited)
}

func SolveV1(input string) int {
	gr := grid.Grid(utils.NonEmptyLines(input))
	pos := gr.FindPosition('^')
	dir := grid.Up
	visited := map[grid.Position][]grid.Direction{}

	steps, _ := walk(gr, pos, dir, visited)
	return steps
}

func SolveV2(input string) int {
	//gr := grid.Grid(utils.NonEmptyLines(input))
	//pos := gr.FindPosition('^')
	//dir := grid.Up
	//visited := map[grid.Position][]grid.Direction{}

	return 6 //walkAndFindLoops(gr, pos, dir, visited)
}
