package day22

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9
`

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 5, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, 7, res)
}
