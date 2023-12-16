package day16

import (
	"fmt"

	"adventofcode.com/internal/utils"
)

type Vector struct{ DX, DY int }

var (
	vecUp    = Vector{-1, 0}
	vecDown  = Vector{1, 0}
	vecRight = Vector{0, 1}
	vecLeft  = Vector{0, -1}
)

type Tracer struct {
	N, M      int
	field     []string
	energized map[string]bool
	traced    map[string]bool
}

func NewTracer(field []string) *Tracer {
	return &Tracer{
		N:         len(field),
		M:         len(field[0]),
		field:     field,
		energized: make(map[string]bool),
		traced:    make(map[string]bool),
	}
}

func (t *Tracer) Beam(i, j int, dir Vector) {
	if i < 0 || t.N <= i || j < 0 || t.M <= j {
		return
	}
	beamLoc := fmt.Sprintf("%d,%d,%d,%d", i, j, dir.DX, dir.DY)
	if t.traced[beamLoc] {
		return
	}
	t.traced[beamLoc] = true
	t.energized[fmt.Sprintf("%d,%d", i, j)] = true

	ch := t.field[i][j]
	if ch == '.' ||
		ch == '-' && (dir == vecLeft || dir == vecRight) ||
		ch == '|' && (dir == vecUp || dir == vecDown) {

		t.Beam(i+dir.DX, j+dir.DY, dir)
	}

	if ch == '-' && (dir == vecUp || dir == vecDown) ||
		ch == '\\' && dir == vecUp ||
		ch == '/' && dir == vecDown {

		t.Beam(i+vecLeft.DX, j+vecLeft.DY, vecLeft)
	}

	if ch == '-' && (dir == vecUp || dir == vecDown) ||
		ch == '\\' && dir == vecDown ||
		ch == '/' && dir == vecUp {

		t.Beam(i+vecRight.DX, j+vecRight.DY, vecRight)
	}

	if ch == '|' && (dir == vecLeft || dir == vecRight) ||
		ch == '\\' && dir == vecLeft ||
		ch == '/' && dir == vecRight {

		t.Beam(i+vecUp.DX, j+vecUp.DY, vecUp)
	}

	if ch == '|' && (dir == vecLeft || dir == vecRight) ||
		ch == '\\' && dir == vecRight ||
		ch == '/' && dir == vecLeft {

		t.Beam(i+vecDown.DX, j+vecDown.DY, vecDown)
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
	field := utils.NonEmptyLines(input)
	t := NewTracer(field)
	t.Beam(0, 0, vecRight)
	return t.Energized()
}

func SolveV2(input string) int {
	field := utils.NonEmptyLines(input)
	res := 0
	t := NewTracer(field)

	for i := 0; i < t.N; i++ {
		t.Beam(i, 0, vecRight)
		res = max(res, t.Energized())
		t.Clear()

		t.Beam(i, t.M-1, vecLeft)
		res = max(res, t.Energized())
		t.Clear()
	}

	for j := 0; j < t.M; j++ {
		t.Beam(0, j, vecDown)
		res = max(res, t.Energized())
		t.Clear()

		t.Beam(t.N-1, j, vecUp)
		res = max(res, t.Energized())
		t.Clear()
	}

	return res
}
