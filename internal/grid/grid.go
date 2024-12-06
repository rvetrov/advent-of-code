package grid

import (
	"fmt"
	"strings"
)

type Grid struct {
	rows int
	cols int
	g    [][]byte
}

func New(lines []string) Grid {
	var bs [][]byte
	for _, line := range lines {
		bs = append(bs, []byte(line))
	}
	return NewFromBytes(bs)
}

func NewFromBytes(bs [][]byte) Grid {
	return Grid{
		rows: len(bs),
		cols: len(bs[0]),
		g:    bs,
	}
}

func (g Grid) Rows() int {
	return g.rows
}

func (g Grid) Cols() int {
	return g.cols
}

func (g Grid) Lines() []string {
	var lines []string
	for _, bytes := range g.g {
		lines = append(lines, string(bytes))
	}
	return lines
}

func (g Grid) Contains(p Position) bool {
	return 0 <= p.Row && p.Row < g.rows && 0 <= p.Col && p.Col < g.cols
}

func (g Grid) At(p Position) (byte, bool) {
	if !g.Contains(p) {
		return 0, false
	}
	return g.g[p.Row][p.Col], true
}

func (g Grid) SetAt(p Position, ch byte) bool {
	if !g.Contains(p) {
		return false
	}
	g.g[p.Row][p.Col] = ch
	return true
}

func (g Grid) FindPosition(ch byte) (Position, bool) {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			if g.g[i][j] == ch {
				return Position{Row: i, Col: j}, true
			}
		}
	}
	return Position{}, false
}

func (g Grid) Start() Position {
	return Position{}
}

func (g Grid) Next(pos Position) Position {
	pos.Col += 1
	if pos.Col >= g.cols {
		pos.Col = 0
		pos.Row += 1
	}
	return pos
}

func (g Grid) EncodePosition(p Position) int {
	return p.Row*g.cols + p.Col
}

func (g Grid) DecodePosition(v int) Position {
	return Position{Row: v / g.cols, Col: v % g.cols}
}

func Print(grid Grid) {
	for _, line := range grid.g {
		fmt.Println(line)
	}
}

func buildersToLines(builders []strings.Builder) Grid {
	res := make([]string, len(builders))
	for i, builder := range builders {
		res[i] = builder.String()
	}
	return New(res)
}

func Transpose(grid Grid) Grid {
	builders := make([]strings.Builder, grid.cols)

	for _, bytes := range grid.g {
		for j, ch := range bytes {
			builders[j].WriteByte(ch)
		}
	}
	return buildersToLines(builders)
}

func RotateCW(grid Grid) Grid {
	builders := make([]strings.Builder, grid.cols)
	for j := 0; j < grid.cols; j++ {
		builder := &builders[j]
		for i := grid.rows - 1; i >= 0; i-- {
			builder.WriteByte(grid.g[i][j])
		}
	}
	return buildersToLines(builders)
}

func RotateCCW(grid Grid) Grid {
	n, m := grid.rows, grid.cols
	builders := make([]strings.Builder, m)
	for j := 0; j < m; j++ {
		builder := &builders[m-1-j]
		for i := 0; i < n; i++ {
			builder.WriteByte(grid.g[i][j])
		}
	}
	return buildersToLines(builders)
}
