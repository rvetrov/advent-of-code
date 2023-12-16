package day07

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

const testCase2 = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 6440, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase2)
	require.Equal(t, 5905, res)
}
