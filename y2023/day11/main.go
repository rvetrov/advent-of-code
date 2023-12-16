package day11

import (
	"strings"

	"adventofcode.com/internal/math"
	"adventofcode.com/internal/utils"
)

type Point struct{ X, Y int }

type Field struct {
	grid      []string
	n         int
	m         int
	emptyRows []bool
	emptyCols []bool
	Galaxies  []Point
}

func countEmpty(i1, i2 int, empty []bool) int {
	if i1 > i2 {
		i1, i2 = i2, i1
	}
	res := 0
	for i := i1; i <= i2; i++ {
		if empty[i] {
			res++
		}
	}
	return res
}

func (f Field) Dist(g1, g2 Point, expansion int) int {
	res := math.AbsInt(g1.X-g2.X) + math.AbsInt(g1.Y-g2.Y)
	res += countEmpty(g1.X, g2.X, f.emptyRows) * (expansion - 1)
	res += countEmpty(g1.Y, g2.Y, f.emptyCols) * (expansion - 1)
	return res
}

func parseField(input string) Field {
	f := Field{grid: []string{}}
	lines := utils.NonEmptyLines(input)
	for _, line := range lines {
		f.grid = append(f.grid, line)
		f.emptyRows = append(f.emptyRows, strings.Count(line, ".") == len(line))
		for j, ch := range line {
			if ch == '#' {
				f.Galaxies = append(f.Galaxies, Point{len(f.grid) - 1, j})
			}
		}
	}
	f.n = len(f.grid)
	f.m = len(f.grid[0])

	for j := 0; j < f.m; j++ {
		dotCnt := 0
		for i := 0; i < f.n; i++ {
			if f.grid[i][j] == '.' {
				dotCnt++
			}
		}
		f.emptyCols = append(f.emptyCols, dotCnt == f.m)
	}
	return f
}

func solve(input string, expansion int) int {
	f := parseField(input)
	res := 0
	for i, g1 := range f.Galaxies {
		for j := i + 1; j < len(f.Galaxies); j++ {
			g2 := f.Galaxies[j]
			res += f.Dist(g1, g2, expansion)
		}
	}
	return res
}

func SolveV1(input string) int {
	return solve(input, 2)
}

func SolveV2(input string) int {
	return solve(input, 1000000)
}
