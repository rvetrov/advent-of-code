package day17

import (
	"container/heap"
	"fmt"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

var rotations = map[string][]grid.Direction{
	grid.Up.String():    {grid.Left, grid.Right},
	grid.Down.String():  {grid.Left, grid.Right},
	grid.Left.String():  {grid.Up, grid.Down},
	grid.Right.String(): {grid.Up, grid.Down},
}

type Crucible interface {
	Moves(MoveState) []MoveState
	CanStop(MoveState) bool
}

type regularCrucible struct{}

func (rc regularCrucible) Moves(state MoveState) []MoveState {
	res := []MoveState{}
	if state.Steps < 3 {
		res = append(res, MoveState{
			Pos:   state.Pos.Add(state.Dir),
			Dir:   state.Dir,
			Steps: state.Steps + 1,
		})
	}
	for _, dir := range rotations[state.Dir.String()] {
		res = append(res, MoveState{
			Pos:   state.Pos.Add(dir),
			Dir:   dir,
			Steps: 1,
		})
	}
	return res
}

func (rc regularCrucible) CanStop(_ MoveState) bool {
	return true
}

type ultraCrucible struct{}

func (uc ultraCrucible) Moves(state MoveState) []MoveState {
	res := []MoveState{}
	if state.Steps < 10 {
		res = append(res, MoveState{
			Pos:   state.Pos.Add(state.Dir),
			Dir:   state.Dir,
			Steps: state.Steps + 1,
		})
	}
	if state.Steps == 0 || state.Steps >= 4 {
		for _, dir := range rotations[state.Dir.String()] {
			res = append(res, MoveState{
				Pos:   state.Pos.Add(dir),
				Dir:   dir,
				Steps: 1,
			})
		}
	}
	return res
}

func (uc ultraCrucible) CanStop(state MoveState) bool {
	return state.Steps >= 4
}

type Solver struct {
	grid       grid.Grid
	distToPos  map[string]int
	distToCell map[string]int
	q          *Heap
}

func NewSolver(gr grid.Grid) *Solver {
	return &Solver{
		grid:       gr,
		distToPos:  make(map[string]int),
		distToCell: make(map[string]int),
		q:          &Heap{ds: make([]Distance, 0)},
	}
}

func (s *Solver) visitPosition(state MoveState, dist int, mover Crucible) {
	s.distToPos[state.String()] = dist
	cellPos := state.Pos.String()
	if mover.CanStop(state) {
		if curDist, found := s.distToCell[cellPos]; !found || curDist > dist {
			s.distToCell[cellPos] = dist
		}
	}
}

func (s *Solver) heatLoss(p grid.Position) int {
	if ch, ok := s.grid.At(p); ok {
		return int(ch) - int('0')
	}
	return 0
}

func (s *Solver) Walk(startPos MoveState, mover Crucible) {
	heap.Push(s.q, Distance{State: startPos, Dist: 0})

	for s.q.Len() > 0 {
		d := heap.Pop(s.q).(Distance)

		if !s.grid.Contains(d.State.Pos) {
			continue
		}
		if savedDist, found := s.distToPos[d.State.String()]; found && savedDist <= d.Dist {
			continue
		}
		s.visitPosition(d.State, d.Dist, mover)

		for _, state := range mover.Moves(d.State) {
			newD := Distance{
				State: state,
				Dist:  d.Dist + s.heatLoss(state.Pos),
			}
			heap.Push(s.q, newD)
		}
	}
}

func (s *Solver) DistToCell(p grid.Position) int {
	return s.distToCell[p.String()]
}

type MoveState struct {
	Pos   grid.Position
	Dir   grid.Direction
	Steps int
}

func (s MoveState) String() string {
	return fmt.Sprintf("%v,%v->%d", s.Pos, s.Dir, s.Steps)
}

type Distance struct {
	State MoveState
	Dist  int
}

type Heap struct {
	ds []Distance
}

func (h *Heap) Len() int           { return len(h.ds) }
func (h *Heap) Less(i, j int) bool { return h.ds[i].Dist < h.ds[j].Dist }
func (h *Heap) Swap(i, j int)      { h.ds[i], h.ds[j] = h.ds[j], h.ds[i] }
func (h *Heap) Push(d any)         { h.ds = append(h.ds, d.(Distance)) }
func (h *Heap) Pop() any {
	d := h.ds[len(h.ds)-1]
	h.ds = h.ds[:len(h.ds)-1]
	return d
}

func solve(input string, mover Crucible) int {
	gr := grid.New(utils.NonEmptyLines(input))

	st := NewSolver(gr)
	startPos := MoveState{Pos: grid.Position{Row: 0, Col: 0}, Dir: grid.Right, Steps: 0}
	st.Walk(startPos, mover)

	return st.DistToCell(grid.Position{Row: gr.Rows() - 1, Col: gr.Cols() - 1})
}

func SolveV1(input string) int {
	return solve(input, regularCrucible{})
}

func SolveV2(input string) int {
	return solve(input, ultraCrucible{})
}
