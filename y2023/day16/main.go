package day16

import (
	"fmt"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

type Tracer struct {
	N, M      int
	grid      []string
	energized map[string]bool
	traced    map[string]bool
}

func NewTracer(gr grid.Grid) *Tracer {
	return &Tracer{
		N:         len(gr),
		M:         len(gr[0]),
		grid:      gr,
		energized: make(map[string]bool),
		traced:    make(map[string]bool),
	}
}

func (t *Tracer) Beam(i, j int, dir grid.Direction) {
	if i < 0 || t.N <= i || j < 0 || t.M <= j {
		return
	}
	beamLoc := fmt.Sprintf("%d,%d,%d,%d", i, j, dir.DR, dir.DC)
	if t.traced[beamLoc] {
		return
	}
	t.traced[beamLoc] = true
	t.energized[fmt.Sprintf("%d,%d", i, j)] = true

	ch := t.grid[i][j]
	if ch == '.' ||
		ch == '-' && (dir == grid.Left || dir == grid.Right) ||
		ch == '|' && (dir == grid.Up || dir == grid.Down) {

		t.Beam(i+dir.DR, j+dir.DC, dir)
	}

	if ch == '-' && (dir == grid.Up || dir == grid.Down) ||
		ch == '\\' && dir == grid.Up ||
		ch == '/' && dir == grid.Down {

		t.Beam(i+grid.Left.DR, j+grid.Left.DC, grid.Left)
	}

	if ch == '-' && (dir == grid.Up || dir == grid.Down) ||
		ch == '\\' && dir == grid.Down ||
		ch == '/' && dir == grid.Up {

		t.Beam(i+grid.Right.DR, j+grid.Right.DC, grid.Right)
	}

	if ch == '|' && (dir == grid.Left || dir == grid.Right) ||
		ch == '\\' && dir == grid.Left ||
		ch == '/' && dir == grid.Right {

		t.Beam(i+grid.Up.DR, j+grid.Up.DC, grid.Up)
	}

	if ch == '|' && (dir == grid.Left || dir == grid.Right) ||
		ch == '\\' && dir == grid.Right ||
		ch == '/' && dir == grid.Left {

		t.Beam(i+grid.Down.DR, j+grid.Down.DC, grid.Down)
	}
}

func (t *Tracer) Energized() int {
	return len(t.energized)
}

func (t *Tracer) Clear() {
	t.energized = make(map[string]bool)
	t.traced = make(map[string]bool)
}

func SolveV1(input string) int {
	gr := utils.NonEmptyLines(input)
	t := NewTracer(gr)
	t.Beam(0, 0, grid.Right)
	return t.Energized()
}

func SolveV2(input string) int {
	gr := utils.NonEmptyLines(input)
	res := 0
	t := NewTracer(gr)

	for i := 0; i < t.N; i++ {
		t.Beam(i, 0, grid.Right)
		res = max(res, t.Energized())
		t.Clear()

		t.Beam(i, t.M-1, grid.Left)
		res = max(res, t.Energized())
		t.Clear()
	}

	for j := 0; j < t.M; j++ {
		t.Beam(0, j, grid.Down)
		res = max(res, t.Energized())
		t.Clear()

		t.Beam(t.N-1, j, grid.Up)
		res = max(res, t.Energized())
		t.Clear()
	}

	return res
}
