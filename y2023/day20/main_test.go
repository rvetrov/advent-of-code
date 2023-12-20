package day20

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`

const testCase2 = `
broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> rx
`

func TestV1(t *testing.T) {
	res := solve(testCase1, 1)
	require.Equal(t, 32, res)

	res = SolveV1(testCase1)
	require.Equal(t, 32000000, res)
	res = SolveV1(testCase2)
	require.Equal(t, 11687500, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase2)
	require.Equal(t, 1, res)
}
