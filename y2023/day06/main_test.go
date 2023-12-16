package day06

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
Time:      7  15   30
Distance:  9  40  200
`

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 288, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, 71503, res)
}
