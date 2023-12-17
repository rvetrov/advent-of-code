package grid

import (
	"fmt"
	"strings"
)

type Grid []string

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
