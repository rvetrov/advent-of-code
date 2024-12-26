package day24

import (
	"testing"

	"github.com/stretchr/testify/require"
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
	res := SolveV1(testCase1)
	require.Equal(t, 0, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, 0, res)
}
