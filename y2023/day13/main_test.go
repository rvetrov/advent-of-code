package day13

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`

const testCase2 = `
#..##...#
#..#..###
#..#..###
#..##...#
.##..#.#.
.....###.
#..#.##.#
...#..###
.##.#..##
.##.#.#.#
#####.#..
`

const testCase3 = `
#..##.#
#.###..
#.###..
#..##.#
.##.##.
...#...
#..###.
#..##..
#..#..#
#..#..#
#..##..
#..###.
.#.#...
`

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 405, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, 400, res)

	res = SolveV2(testCase2)
	require.Equal(t, 2, res)

	res = SolveV2(testCase3)
	require.Equal(t, 900, res)
}
