package day05

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase0 = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase0)
	require.Equal(t, 143, res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase0)
	require.Equal(t, 123, res)
}
