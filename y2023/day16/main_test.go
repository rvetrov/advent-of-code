package day16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
`

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 46, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, 51, res)
}
