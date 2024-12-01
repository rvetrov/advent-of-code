package day18

import (
	"fmt"
	"log"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

type Dig struct {
	dir   string
	Dist  int
	Color string
}

func (d Dig) Dir() grid.Direction {
	return directions[d.dir]
}

func (d Dig) ExtractColor() Dig {
	res := Dig{Color: d.Color}
	fmt.Sscanf(d.Color[:5], "%x", &res.Dist)
	switch d.Color[5] {
	case '0':
		res.dir = "R"
	case '1':
		res.dir = "D"
	case '2':
		res.dir = "L"
	case '3':
		res.dir = "U"
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
	res := []Dig{}
	for _, line := range utils.NonEmptyLines(input) {
		d := Dig{}
		if cnt, err := fmt.Sscanf(line, "%s %d (#%6s)", &d.dir, &d.Dist, &d.Color); err == nil && cnt == 3 {
			res = append(res, d)
		} else {
			log.Fatalf("Line: %q, Err: %q, Count: %d, Parsed: %v", line, err, cnt, d)
		}
	}
	return res
}

func buildGrid(digs []Dig) grid.Grid {
	cur := grid.Position{}
	var mx, mn grid.Position
	for _, dig := range digs {
		cur = cur.Add(dig.Dir().Multiplied(dig.Dist))
		mx.Row = max(mx.Row, cur.Row)
		mx.Col = max(mx.Col, cur.Col)
		mn.Row = min(mn.Row, cur.Row)
		mn.Col = min(mn.Col, cur.Col)
	}

	n := mx.Row - mn.Row + 1
	m := mx.Col - mn.Col + 1
	gr := make([][]byte, n)
	for i := 0; i < n; i++ {
		gr[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			gr[i][j] = '.'
		}
	}

	cur = grid.Position{}.Subtract(mn).AsPosition()
	for _, dig := range digs {
		gr[cur.Row][cur.Col] = '#'
		for i := 0; i < dig.Dist; i++ {
			cur = cur.Add(dig.Dir())
			gr[cur.Row][cur.Col] = '#'
		}
	}

	res := make(grid.Grid, n)
	for i, line := range gr {
		res[i] = string(line)
	}
	return res
}

func examineExterior(pos grid.Position, gr grid.Grid, visited []bool) int {
	if !gr.Contains(pos) || gr[pos.Row][pos.Col] != '.' || visited[pos.Row*gr.Cols()+pos.Col] {
		return 0
	}
	visited[pos.Row*gr.Cols()+pos.Col] = true
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
	//grid.Print(gr)

	res := lagoonVolume(gr)
	return res
}

func SolveV2(input string) int {
	// digs := parseDigs(input)
	// for i, dig := range digs {
	// 	digs[i] = dig.ExtractColor()
	// 	fmt.Println(digs[i])
	// }
	// gr := buildGrid(digs)
	// fmt.Println(gr.Rows(), gr.Cols())
	return 0
}
