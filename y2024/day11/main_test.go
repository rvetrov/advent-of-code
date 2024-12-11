package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase0 = `125 17`
)

func TestBlink(t *testing.T) {
	for _, tc := range []struct {
		chain []string
	}{
		{[]string{
			"0 1 10 99 999",
			"1 2024 1 0 9 9 2021976",
		}},
		{[]string{
			"125 17",
			"253000 1 7",
			"253 0 2024 14168",
			"512072 1 20 24 28676032",
			"512 72 2024 2 0 2 4 2867 6032",
			"1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32",
			"2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2",
		}},
	} {
		for i := 1; i < len(tc.chain); i++ {
			before := readStones(tc.chain[i-1])
			after := blink(before)
			expected := readStones(tc.chain[i])

			require.Equal(t, expected, after)
		}
	}
}

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase0)
	require.Equal(t, 55312, res)
}
