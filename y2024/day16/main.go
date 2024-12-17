package day16

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
	"container/heap"
)

type ReindeerPosition struct {
	pos grid.Position
	dir grid.Direction
}

type ReindeerState struct {
	pos  ReindeerPosition
	dist int
}

type Frontier struct {
	states []ReindeerState
}

func (f *Frontier) Len() int {
	return len(f.states)
}

func (f *Frontier) Less(i, j int) bool {
	return f.states[i].dist < f.states[j].dist
}

func (f *Frontier) Swap(i, j int) {
	f.states[i], f.states[j] = f.states[j], f.states[i]
}

func (f *Frontier) Push(x any) {
	pos := x.(ReindeerState)
	f.states = append(f.states, pos)
}

func (f *Frontier) Pop() any {
	last := f.states[len(f.states)-1]
	f.states = f.states[:len(f.states)-1]
	return last
}

func findRoute(gr grid.Grid, start, end grid.Position, dist map[ReindeerPosition]int) int {
	frontier := &Frontier{}
	startState := ReindeerState{pos: ReindeerPosition{pos: start, dir: grid.Right}, dist: 0}
	heap.Push(frontier, startState)

	for frontier.Len() > 0 {
		state := heap.Pop(frontier).(ReindeerState)
		pos := state.pos
		if _, ok := dist[pos]; ok {
			continue
		}
		dist[pos] = state.dist

		if pos.pos == end {
			return state.dist
		}

		for _, nextState := range []ReindeerState{
			{pos: ReindeerPosition{pos: pos.pos.Add(pos.dir), dir: pos.dir}, dist: state.dist + 1},
			{pos: ReindeerPosition{pos: pos.pos, dir: pos.dir.TurnCW()}, dist: state.dist + 1000},
			{pos: ReindeerPosition{pos: pos.pos, dir: pos.dir.TurnCCW()}, dist: state.dist + 1000},
		} {
			if ch, inside := gr.At(nextState.pos.pos); inside && ch != '#' {
				if _, ok := dist[nextState.pos]; !ok {
					heap.Push(frontier, nextState)
				}
			}
		}
	}
	return 0
}

func backwardRoute(gr grid.Grid, pos ReindeerPosition, dist map[ReindeerPosition]int, visited map[grid.Position]struct{}) {
	visited[pos.pos] = struct{}{}
	d := dist[pos]
	delete(dist, pos)
	if d == 0 {
		return
	}

	for _, prevState := range []ReindeerState{
		{pos: ReindeerPosition{pos: pos.pos.Add(pos.dir.Reversed()), dir: pos.dir}, dist: d - 1},
		{pos: ReindeerPosition{pos: pos.pos, dir: pos.dir.TurnCW()}, dist: d - 1000},
		{pos: ReindeerPosition{pos: pos.pos, dir: pos.dir.TurnCCW()}, dist: d - 1000},
	} {
		if ch, inside := gr.At(prevState.pos.pos); inside && ch != '#' {
			if prevDist, ok := dist[prevState.pos]; ok && prevDist == prevState.dist {
				backwardRoute(gr, prevState.pos, dist, visited)
			}
		}
	}
}

func SolveV1(input string) int {
	gr := grid.New(utils.NonEmptyLines(input))
	start, _ := gr.FindPosition('S')
	end, _ := gr.FindPosition('E')
	dist := make(map[ReindeerPosition]int)
	return findRoute(gr, start, end, dist)
}

func SolveV2(input string) int {
	gr := grid.New(utils.NonEmptyLines(input))
	start, _ := gr.FindPosition('S')
	end, _ := gr.FindPosition('E')
	dist := make(map[ReindeerPosition]int)
	resultDist := findRoute(gr, start, end, dist)

	visited := make(map[grid.Position]struct{})
	for _, finalDir := range grid.FourSides {
		finalPos := ReindeerPosition{pos: end, dir: finalDir}
		if d, ok := dist[finalPos]; ok && d == resultDist {
			backwardRoute(gr, finalPos, dist, visited)
		}
	}
	return len(visited)
}
