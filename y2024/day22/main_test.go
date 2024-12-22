package day22

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase1 = `
1
10
100
2024
`
	testCase2 = `
1
2
3
2024
`
)

func TestSolveV1(t *testing.T) {
	for _, tc := range []struct{ num, expected int }{
		{1, 8685429},
		{10, 4700978},
		{100, 15273692},
		{2024, 8667524},
	} {
		got := simulateRandom(tc.num, 2000)
		require.Equal(t, tc.expected, got)
	}

	res := SolveV1(testCase1)
	require.Equal(t, 37327623, res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase2)
	require.Equal(t, 23, res)
}
