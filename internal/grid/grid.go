package grid

import (
	"fmt"
	"strings"
)

type Grid []string

func (g Grid) Rows() int {
	return len(g)
}

func (g Grid) Cols() int {
	if len(g) == 0 {
		return 0
	}
	return len(g[0])
}

func (g Grid) Contains(p Position) bool {
	return 0 <= p.Row && p.Row < len(g) && 0 <= p.Col && p.Col < len(g[p.Row])
}

func (g Grid) EncodePosition(p Position) int {
	return p.Row*g.Cols() + p.Col
}

func (g Grid) DecodePosition(v int) Position {
	rowSize := g.Cols()
	return Position{Row: v / rowSize, Col: v % rowSize}
}

func Print(grid Grid) {
	for _, line := range grid {
		fmt.Println(line)
	}
}

func buildersToLines(builders []strings.Builder) Grid {
	res := make([]string, len(builders))
	for i, builder := range builders {
		res[i] = builder.String()
	}
	return res
}

func Transpose(grid Grid) Grid {
	n, m := len(grid), len(grid[0])
	builders := make([]strings.Builder, m)

	for i := 0; i < n; i++ {
		for j, ch := range grid[i] {
			builders[j].WriteRune(ch)
		}
	}
	return buildersToLines(builders)
}

func RotateCW(grid Grid) Grid {
	n, m := len(grid), len(grid[0])
	builders := make([]strings.Builder, m)
	for j := 0; j < m; j++ {
		builder := &builders[j]
		for i := n - 1; i >= 0; i-- {
			builder.WriteByte(grid[i][j])
		}
	}
	return buildersToLines(builders)
}

func RotateCCW(grid Grid) Grid {
	n, m := len(grid), len(grid[0])
	builders := make([]strings.Builder, m)
	for j := 0; j < m; j++ {
		builder := &builders[m-1-j]
		for i := 0; i < n; i++ {
			builder.WriteByte(grid[i][j])
		}
	}
	return buildersToLines(builders)
}
