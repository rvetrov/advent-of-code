package day17

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
`

const testCase2 = `
111111111111
999999999991
999999999991
999999999991
999999999991
`

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 102, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, 94, res)

	res = SolveV2(testCase2)
	require.Equal(t, 71, res)
}
