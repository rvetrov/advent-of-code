package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

func TestV1(t *testing.T) {
	res := solveV1(testCase1)
	require.Equal(t, 374, res)
}

func TestV2(t *testing.T) {
	res := solve(testCase1, 10)
	require.Equal(t, 1030, res)

	res = solve(testCase1, 100)
	require.Equal(t, 8410, res)
}
