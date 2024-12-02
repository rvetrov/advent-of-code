package day02

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestSolveV1(t *testing.T) {
	res1 := SolveV1(testCase1)
	require.Equal(t, 2, res1)
}

func TestSolveV2(t *testing.T) {
	res1 := SolveV2(testCase1)
	require.Equal(t, 4, res1)
}

func TestIsSafeReport(t *testing.T) {
	for _, tc := range []struct {
		nums          []int
		allowBadLevel bool
		safe          bool
	}{
		{[]int{5, 8, 9}, true, true},
		{[]int{5, 0, 4}, true, true},
		{[]int{5, 0, 6}, true, true},
		{[]int{7, 6, 4, 2, 1}, true, true},
		{[]int{1, 2, 7, 8, 9}, true, false},
		{[]int{9, 7, 6, 2, 1}, true, false},
		{[]int{1, 3, 2, 4, 5}, true, true},
		{[]int{8, 6, 4, 4, 1}, true, true},
		{[]int{1, 3, 6, 7, 9}, true, true},
	} {
		require.Equal(t, tc.safe, isSafeReport(tc.nums, tc.allowBadLevel), tc)
	}
}
