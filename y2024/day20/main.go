package day20

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

const (
	wallChar = '#'
)

func distsOnRoute(gr grid.Grid, start grid.Position) map[grid.Position]int {
	dists := make(map[grid.Position]int)

	q := []grid.Position{start}
	dists[start] = 0

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

func cheatRoutes(gr grid.Grid, dists map[grid.Position]int, start grid.Position, steps int, threshold int) int {
	visited := make(map[grid.Position]struct{})
	visited[start] = struct{}{}

	cheatStartDist := dists[start]
	cheatEnds := make(map[grid.Position]int)

	wave := []grid.Position{start}
	for step := range steps {
		var newWave []grid.Position

		for _, pos := range wave {
			for _, dir := range grid.FourSides {
				nextPos := pos.Add(dir)
				if _, inside := gr.At(nextPos); !inside {
					continue
				}
				if _, v := visited[nextPos]; v {
					continue
				}
				visited[nextPos] = struct{}{}

				if cheatEndDist, onRoute := dists[nextPos]; onRoute {
					saved := cheatEndDist - cheatStartDist - (step + 1)
					if saved > 0 {
						cheatEnds[nextPos] = saved
					}
				}
				newWave = append(newWave, nextPos)
			}
		}

		wave = newWave
	}

	res := 0
	for _, saved := range cheatEnds {
		if saved >= threshold {
			res++
		}
	}
	return res
}

func solve(input string, threshold int, cheatSteps int) int {
	gr := grid.New(utils.NonEmptyLines(input))
	start, _ := gr.FindPosition('S')
	dists := distsOnRoute(gr, start)

	res := 0
	for cheatStart := range dists {
		res += cheatRoutes(gr, dists, cheatStart, cheatSteps, threshold)
	}
	return res
}

func SolveV1(input string) int {
	return solve(input, 100, 2)
}

func SolveV2(input string) int {
	return solve(input, 100, 20)
}
