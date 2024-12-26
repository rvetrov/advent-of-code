package day23

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

type Distances map[int]map[int]int

const (
	blockedCh = '#'
	visitedCh = 'O'
	notFound  = -1
)

var (
	slopeToDir = map[byte]grid.Direction{
		'<': grid.Left,
		'>': grid.Right,
		'v': grid.Down,
		'^': grid.Up,
	}
)

func dfs1(pos grid.Position, gr grid.Grid) (int, bool) {
	origCh, inside := gr.At(pos)
	if !inside {
		return 0, true
	} else if origCh == visitedCh || origCh == blockedCh {
		return 0, false
	}

	gr.SetAt(pos, visitedCh)
	defer gr.SetAt(pos, origCh)

	if slopeToDir != nil {
		slopeDir, isSlope := slopeToDir[origCh]
		if isSlope {
			mx, found := dfs1(pos.Add(slopeDir), gr)
			return mx + 1, found
		}
	}

	mx, found := 0, false
	for _, dir := range grid.FourSides {
		newPos := pos.Add(dir)
		newMx, newFound := dfs1(newPos, gr)
		if newFound && (!found || mx < newMx) {
			mx = newMx
			found = true
		}
	}
	return mx + 1, found
}

func SolveV1(input string) int {
	gr := grid.New(utils.NonEmptyLines(input))
	steps, _ := dfs1(grid.Position{Col: 1}, gr)
	return steps - 1
}

func isGraphNode(gr grid.Grid, pos grid.Position) bool {
	if pos.Row == 0 || pos.Col == 0 || pos.Row == gr.Rows()-1 || pos.Col == gr.Cols()-1 {
		return true
	}
	ns := 0
	for _, dir := range grid.FourSides {
		if ch, inside := gr.At(pos.Add(dir)); inside && ch != blockedCh {
			ns++
		}
	}
	return ns > 2
}

func closestGraphNode(gr grid.Grid, pos grid.Position) (int, int) {
	if isGraphNode(gr, pos) {
		return gr.EncodePosition(pos), 0
	}

	origCh, _ := gr.At(pos)
	gr.SetAt(pos, visitedCh)
	defer gr.SetAt(pos, origCh)

	for _, dir := range grid.FourSides {
		newPos := pos.Add(dir)
		if newCh, inside := gr.At(newPos); inside && newCh != visitedCh && newCh != blockedCh {
			node, dist := closestGraphNode(gr, newPos)
			return node, dist + 1
		}
	}
	return notFound, 0
}

func closestGraphNodes(gr grid.Grid, pos grid.Position) map[int]int {
	origCh, _ := gr.At(pos)
	gr.SetAt(pos, visitedCh)
	defer gr.SetAt(pos, origCh)

	closestNodes := make(map[int]int)
	for _, dir := range grid.FourSides {
		startPos := pos.Add(dir)
		node, dist := closestGraphNode(gr, startPos)
		if node != notFound {
			if prev, exists := closestNodes[node]; !exists || prev < dist+1 {
				closestNodes[node] = dist + 1
			}
		}
	}
	return closestNodes
}

func dfs2(pos, finish int, dists Distances, visited map[int]struct{}) (int, bool) {
	if pos == finish {
		return 0, true
	}
	visited[pos] = struct{}{}
	defer delete(visited, pos)

	mx, found := 0, false
	for neighbour, dist := range dists[pos] {
		if _, v := visited[neighbour]; v {
			continue
		}
		newMx, newFound := dfs2(neighbour, finish, dists, visited)
		if newFound && (!found || mx < newMx+dist) {
			mx = newMx + dist
			found = true
		}
	}
	return mx, found
}

func SolveV2(input string) int {
	var start, finish int
	dists := make(Distances)
	gr := grid.New(utils.NonEmptyLines(input))

	for pos := gr.First(); gr.Contains(pos); pos = gr.Next(pos) {
		if ch, _ := gr.At(pos); ch == blockedCh {
			continue
		}
		if isGraphNode(gr, pos) {
			if pos.Row == 0 {
				start = gr.EncodePosition(pos)
			} else if pos.Row == gr.Rows()-1 {
				finish = gr.EncodePosition(pos)
			}
			dists[gr.EncodePosition(pos)] = closestGraphNodes(gr, pos)
		}
	}

	steps, _ := dfs2(start, finish, dists, make(map[int]struct{}))
	return steps
}
