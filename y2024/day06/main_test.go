package day06

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase0 = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase0)
	require.Equal(t, 41, res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase0)
	require.Equal(t, 6, res)
}
