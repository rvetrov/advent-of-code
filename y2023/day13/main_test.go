package main

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
	res := solveV1(testCase1)
	require.Equal(t, 405, res)
}

func TestV2(t *testing.T) {
	res := solveV2(testCase1)
	require.Equal(t, 400, res)

	res = solveV2(testCase2)
	require.Equal(t, 2, res)

	res = solveV2(testCase3)
	require.Equal(t, 900, res)
}
