package day04

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase0 = `
..X...
.SAMX.
.A..A.
XMAS.S
.X....
`

	testCase1 = `
....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
`

	testCase2 = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

	testCase2_0 = `
M.S
.A.
M.S
`
	testCase2_1 = `
.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase0)
	require.Equal(t, 4, res)

	res = SolveV1(testCase1)
	require.Equal(t, 18, res)

	res = SolveV1(testCase2)
	require.Equal(t, 18, res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase2_0)
	require.Equal(t, 1, res)

	res = SolveV2(testCase2_1)
	require.Equal(t, 9, res)
}
