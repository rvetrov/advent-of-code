package day10

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase0 = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase0)
	require.Equal(t, 36, res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase0)
	require.Equal(t, 81, res)
}
