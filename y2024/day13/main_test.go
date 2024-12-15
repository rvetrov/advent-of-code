package day13

import (
	"adventofcode.com/internal/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	testCases = []struct {
		line   string
		result int
	}{
		{
			`
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`,
			280,
		},
		{
			`
Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176`,
			0,
		},
		{
			`
Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450`,
			200,
		},
		{
			`
Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
			0,
		},
		{
			`
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=10000000008400, Y=10000000005400`,
			0,
		},
		{
			`
Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=10000000012748, Y=10000000012176`,
			459236326669,
		},
		{
			`Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=10000000007870, Y=10000000006450`,
			0,
		},
		{
			`Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=10000000018641, Y=10000000010279`,
			416082282239,
		},
	}
)

func TestSolveV1(t *testing.T) {
	for _, tc := range testCases[:4] {
		res := SolveV1(tc.line)
		require.Equal(t, tc.result, res)
	}
}

func TestSolveV2(t *testing.T) {
	for _, tc := range testCases {
		a, b, target := parseInput(utils.NonEmptyLines(tc.line))
		res := solve(a, b, target)
		require.Equal(t, tc.result, res)
	}
}
