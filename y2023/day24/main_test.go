package day24

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase1 = `
19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3
`
)

func TestV1(t *testing.T) {
	res := solveV1(testCase1, 7, 27)
	require.Equal(t, 2, res)
}
