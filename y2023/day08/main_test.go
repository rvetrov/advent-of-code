package day08

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

const testCase2 = `
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

const testCaseV2 = `
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

func TestV1(t *testing.T) {
	res := solveV1(testCase1)
	require.Equal(t, 2, res)

	res = solveV1(testCase2)
	require.Equal(t, 6, res)
}

func TestV2(t *testing.T) {
	res := solveV2(testCaseV2)
	require.Equal(t, 6, res)
}
