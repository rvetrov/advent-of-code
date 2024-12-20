package day20

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase1 = `
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`
)

func TestSolveV1(t *testing.T) {
	expectedCheats := 0
	for _, tc := range []struct{ saved, cheats int }{
		{64, 1},
		{40, 1},
		{38, 1},
		{36, 1},
		{20, 1},
		{12, 3},
		{10, 2},
		{8, 4},
		{6, 2},
		{4, 14},
		{2, 14},
	} {
		expectedCheats += tc.cheats
		res := solve(testCase1, tc.saved, 2)
		require.Equal(t, expectedCheats, res)
	}
}

func TestSolveV2(t *testing.T) {
	expectedCheats := 0
	for ti, tc := range []struct{ saved, cheats int }{
		{76, 3},
		{74, 4},
		{72, 22},
		{70, 12},
		{68, 14},
		{66, 12},
		{64, 19},
		{62, 20},
		{60, 23},
		{58, 25},
		{56, 39},
		{54, 29},
		{52, 31},
		{50, 32},
	} {
		expectedCheats += tc.cheats
		res := solve(testCase1, tc.saved, 20)
		require.Equal(t, expectedCheats, res, ti)
	}
}
