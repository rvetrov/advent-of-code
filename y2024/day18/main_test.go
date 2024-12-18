package day18

import (
	"testing"

	"adventofcode.com/internal/utils"
	"github.com/stretchr/testify/require"
)

const (
	testCase1 = `
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`
)

func TestSolveV1(t *testing.T) {
	gr := makeGrid(7, 7)
	corruptedBytes := utils.NonEmptyLines(testCase1)
	res := solveV1(gr, corruptedBytes[:12])
	require.Equal(t, 22, res)
}

func TestSolveV2(t *testing.T) {
	gr := makeGrid(7, 7)
	corruptedBytes := utils.NonEmptyLines(testCase1)
	res := solveV2(gr, corruptedBytes)
	require.Equal(t, "6,1", res)
}
