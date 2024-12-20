package day20

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/math"
	"adventofcode.com/internal/utils"
)

const (
	startChar = 'S'
	wallChar  = '#'
)

func distsOnRoute(gr grid.Grid) map[grid.Position]int {
	start, _ := gr.FindPosition(startChar)

	dists := make(map[grid.Position]int)
	dists[start] = 0
	q := []grid.Position{start}

	for len(q) > 0 {
		pos := q[0]
		dist, _ := dists[pos]
		q = q[1:]

		for _, dir := range grid.FourSides {
			nextPos := pos.Add(dir)
			if ch, inside := gr.At(nextPos); !inside || ch == wallChar {
				continue
			}
			if _, visited := dists[nextPos]; visited {
				continue
			}
			dists[nextPos] = dist + 1
			q = append(q, nextPos)
		}
	}
	return dists
}

func cheatRoutes(gr grid.Grid, dists map[grid.Position]int, start grid.Position, maxSteps int, minSave int) int {
	cheatStartDist := dists[start]
	res := 0
	for dRow := -maxSteps; dRow <= maxSteps; dRow++ {
		maxColSteps := maxSteps - math.AbsInt(dRow)
		for dCol := -maxColSteps; dCol <= maxColSteps; dCol++ {
			pos := start.Add(grid.Direction{DR: dRow, DC: dCol})
			if _, inside := gr.At(pos); !inside {
				continue
			}

			if cheatEndDist, onRoute := dists[pos]; onRoute {
				steps := math.AbsInt(dRow) + math.AbsInt(dCol)
				saved := cheatEndDist - cheatStartDist - steps
				if saved >= minSave {
					res++
				}
			}
		}
	}
	return res
}

func solve(input string, minSave int, maxSteps int) int {
	gr := grid.New(utils.NonEmptyLines(input))
	dists := distsOnRoute(gr)

	res := 0
	for cheatStart := range dists {
		res += cheatRoutes(gr, dists, cheatStart, maxSteps, minSave)
	}
	return res
}

func SolveV1(input string) int {
	return solve(input, 100, 2)
}

func SolveV2(input string) int {
	return solve(input, 100, 20)
}
