package day19

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase1 = `
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 6, res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, 16, res)
}
