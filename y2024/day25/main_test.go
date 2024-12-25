package day25

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase1 = `
#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 3, res)
}

func TestSolveV2(t *testing.T) {
}
