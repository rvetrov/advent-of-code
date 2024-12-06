package grid

import (
	"fmt"
	"strings"
)

type Grid struct {
	rows int
	cols int
	g    []string
}

func New(lines []string) Grid {
	return Grid{
		rows: len(lines),
		cols: len(lines[0]),
		g:    lines,
	}
}

func NewFromBytes(bs [][]byte) Grid {
	var lines []string
	for _, line := range bs {
		lines = append(lines, string(line))
	}
	return New(lines)
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
		lines = append(lines, bytes)
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

func (g Grid) FindPosition(ch byte) Position {
	n, m := g.Rows(), g.Cols()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if g.g[i][j] == ch {
				return Position{Row: i, Col: j}
			}
		}
	}
	return Position{}
}

func (g Grid) EncodePosition(p Position) int {
	return p.Row*g.Cols() + p.Col
}

func (g Grid) DecodePosition(v int) Position {
	rowSize := g.Cols()
	return Position{Row: v / rowSize, Col: v % rowSize}
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
	n, m := grid.rows, grid.cols
	builders := make([]strings.Builder, m)

	for i := 0; i < n; i++ {
		for j, ch := range grid.g[i] {
			builders[j].WriteRune(ch)
		}
	}
	return buildersToLines(builders)
}

func RotateCW(grid Grid) Grid {
	n, m := grid.rows, grid.cols
	builders := make([]strings.Builder, m)
	for j := 0; j < m; j++ {
		builder := &builders[j]
		for i := n - 1; i >= 0; i-- {
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
