package day09

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase0 = `
2333133121414131402
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase0)
	require.Equal(t, 1928, res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase0)
	require.Equal(t, 2858, res)
}
