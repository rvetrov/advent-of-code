package day21

import (
	"adventofcode.com/internal/grid"
	"testing"

	"github.com/stretchr/testify/require"

	"adventofcode.com/internal/utils"
)

const testCase1 = `
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`

func TestV1(t *testing.T) {
	gr := grid.New(utils.NonEmptyLines(testCase1))
	start := findStart(gr)

	for _, tc := range []struct{ steps, expected int }{
		{1, 2},
		{2, 4},
		{3, 6},
		{6, 16},
	} {
		res := countVisited(gr, start, tc.steps)
		require.Equal(t, tc.expected, res, tc)
	}
}

func TestV2(t *testing.T) {
	gr := grid.New(utils.NonEmptyLines(testCase1))
	start := findStart(gr)

	for _, tc := range []struct{ steps, expected int }{
		{6, 16},
		//{10, 50},
		//{50, 1594},
		//{100, 6536},
		//{500, 167004},
		//{1000, 668697},
		//{5000, 16733044},
	} {
		res := countVisited(gr, start, tc.steps)
		require.Equal(t, tc.expected, res, tc)
	}
}
