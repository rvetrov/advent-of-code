package day03

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	testCase2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
)

func TestSolveV1(t *testing.T) {
	res1 := SolveV1(testCase1)
	require.Equal(t, 161, res1)
}

func TestSolveV2(t *testing.T) {
	res1 := SolveV2(testCase2)
	require.Equal(t, 48, res1)
}

func TestExtractInstructions(t *testing.T) {
	expr := &expression{}
	instructions := expr.extractInstruction(testCase1)
	require.Equal(t, []Instruction{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}, instructions)
}
