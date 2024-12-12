package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase0 = `
AAAA
BBCD
BBCC
EEEC
`
	testCase1 = `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
`
	testCase2 = `
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
`
	testCase3 = `
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
`
	testCase4 = `
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase0)
	require.Equal(t, 140, res)
	res = SolveV1(testCase1)
	require.Equal(t, 1930, res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase0)
	require.Equal(t, 80, res)
	res = SolveV2(testCase1)
	require.Equal(t, 1206, res)
	res = SolveV2(testCase2)
	require.Equal(t, 436, res)
	res = SolveV2(testCase3)
	require.Equal(t, 236, res)
	res = SolveV2(testCase4)
	require.Equal(t, 368, res)
}
