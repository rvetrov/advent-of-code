package day18

import (
	"fmt"
	"log"
	"slices"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
	"golang.org/x/exp/maps"
)

type Dig struct {
	Begin grid.Position
	End   grid.Position
	Dir   grid.Direction
	Dist  int
	Color string
}

func (d Dig) ExtractColor() Dig {
	res := Dig{Color: d.Color}
	_, _ = fmt.Sscanf(d.Color[:5], "%x", &res.Dist)
	switch d.Color[5] {
	case '0':
		res.Dir = grid.Right
	case '1':
		res.Dir = grid.Down
	case '2':
		res.Dir = grid.Left
	case '3':
		res.Dir = grid.Up
	}
	return res
}

var directions = map[string]grid.Direction{
	"U": grid.Up,
	"D": grid.Down,
	"R": grid.Right,
	"L": grid.Left,
}

func parseDigs(input string) []Dig {
	var res []Dig
	var dir string
	for _, line := range utils.NonEmptyLines(input) {
		d := Dig{}
		if cnt, err := fmt.Sscanf(line, "%s %d (#%6s)", &dir, &d.Dist, &d.Color); err == nil && cnt == 3 {
			d.Dir = directions[dir]
			res = append(res, d)
		} else {
			log.Fatalf("Line: %q, Err: %q, Count: %d, Parsed: %v", line, err, cnt, d)
		}
	}
	return res
}

func restoreDigCoordinates(digs []Dig) (mn grid.Position, mx grid.Position) {
	cur := grid.Position{}
	for _, dig := range digs {
		dig.Begin = cur
		cur = cur.Add(dig.Dir.Multiplied(dig.Dist))
		dig.End = cur
		mx.Row = max(mx.Row, cur.Row)
		mx.Col = max(mx.Col, cur.Col)
		mn.Row = min(mn.Row, cur.Row)
		mn.Col = min(mn.Col, cur.Col)
	}
	return mn, mx
}

func buildGrid(digs []Dig) grid.Grid {
	mn, mx := restoreDigCoordinates(digs)

	n := mx.Row - mn.Row + 1
	m := mx.Col - mn.Col + 1
	gr := make([][]byte, n)
	for i := 0; i < n; i++ {
		gr[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			gr[i][j] = '.'
		}
	}

	cur := grid.Position{}.Subtract(mn).AsPosition()
	for _, dig := range digs {
		gr[cur.Row][cur.Col] = '#'
		for i := 0; i < dig.Dist; i++ {
			cur = cur.Add(dig.Dir)
			gr[cur.Row][cur.Col] = '#'
		}
	}

	return grid.NewFromBytes(gr)
}

func examineExterior(pos grid.Position, gr grid.Grid, visited []bool) int {
	if ch, ok := gr.At(pos); !ok || ch != '.' || visited[gr.EncodePosition(pos)] {
		return 0
	}
	visited[gr.EncodePosition(pos)] = true
	res := 1
	for _, dir := range grid.FourSides {
		res += examineExterior(pos.Add(dir), gr, visited)
	}
	return res
}

func lagoonVolume(gr grid.Grid) int {
	n, m := gr.Rows(), gr.Cols()
	visited := make([]bool, n*m)

	exterior := 0
	for i := 0; i < n; i++ {
		exterior += examineExterior(grid.Position{Row: i, Col: 0}, gr, visited)
		exterior += examineExterior(grid.Position{Row: i, Col: m - 1}, gr, visited)
	}
	for j := 0; j < m; j++ {
		exterior += examineExterior(grid.Position{Row: 0, Col: j}, gr, visited)
		exterior += examineExterior(grid.Position{Row: n - 1, Col: j}, gr, visited)
	}
	return n*m - exterior
}

func SolveV1(input string) int {
	digs := parseDigs(input)
	gr := buildGrid(digs)
	return lagoonVolume(gr)
}

func SolveV2(input string) int {
	digs := parseDigs(input)
	for i, dig := range digs {
		digs[i] = dig.ExtractColor()
	}
	restoreDigCoordinates(digs)

	uniqRows := make(map[int]struct{})
	for _, dig := range digs {
		uniqRows[dig.Begin.Row] = struct{}{}
		uniqRows[dig.End.Row] = struct{}{}
	}
	rows := maps.Keys(uniqRows)
	slices.Sort(rows)

	return 0
}
