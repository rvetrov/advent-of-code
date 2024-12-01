package day01

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 string = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 11, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, 31, res)
}
