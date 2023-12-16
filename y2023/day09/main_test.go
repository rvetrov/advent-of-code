package day09

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 114, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, 2, res)
}
