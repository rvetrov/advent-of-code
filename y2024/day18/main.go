package day18

import (
	"strings"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

func makeGrid(rows, cols int) grid.Grid {
	var gridLines []string
	for range rows {
		gridLines = append(gridLines, strings.Repeat(".", cols))
	}
	return grid.New(gridLines)
}

func corruptBytes(gr grid.Grid, corruptedBytes []string) {
	for _, line := range corruptedBytes {
		nums := utils.SplitNumbers(line, ",")
		gr.SetAt(grid.Position{Row: nums[1], Col: nums[0]}, '#')
	}
}

func routeDist(gr grid.Grid) int {
	start, end := gr.First(), gr.Last()
	dist := make(map[grid.Position]int)
	dist[start] = 0
	q := []grid.Position{start}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		curDist := dist[cur]
		for _, dir := range grid.FourSides {
			next := cur.Add(dir)
			if _, ok := dist[next]; ok {
				continue
			}
			if ch, inside := gr.At(next); inside && ch != '#' {
				dist[next] = curDist + 1
				q = append(q, next)
			}
		}
	}
	return dist[end]
}

func solveV1(gr grid.Grid, corruptedBytes []string) int {
	corruptBytes(gr, corruptedBytes)
	return routeDist(gr)
}

func SolveV1(input string) int {
	gr := makeGrid(71, 71)
	corruptedBytes := utils.NonEmptyLines(input)
	return solveV1(gr, corruptedBytes[:1024])
}

func solveV2(rows, cols int, corruptedBytes []string) string {
	i := utils.LowerBound(
		0,
		len(corruptedBytes),
		func(n int) bool {
			gr := makeGrid(rows, cols)
			return solveV1(gr, corruptedBytes[:n+1]) == 0
		},
	)
	return corruptedBytes[i]
}

func SolveV2(input string) string {
	corruptedBytes := utils.NonEmptyLines(input)
	return solveV2(71, 71, corruptedBytes)
}
