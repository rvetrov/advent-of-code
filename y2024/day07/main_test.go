package day07

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase0 = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase0)
	require.Equal(t, 3749, res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase0)
	require.Equal(t, 11387, res)
}
