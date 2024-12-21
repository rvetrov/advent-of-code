package day21

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase1 = `
029A
980A
179A
456A
379A
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 126384, res)
}

func TestSolveV2(t *testing.T) {
}
