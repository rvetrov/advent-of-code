package day17

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	testCase1 = `
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`

	testCase2 = `
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0

A = A >> 3
out A % 8
jmp 0 if A != 0
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, "4,6,3,5,6,3,5,2,1,0", res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase2)
	require.Equal(t, 117440, res)
}
