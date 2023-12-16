package main

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
	res := solveV1(testCase1)
	require.Equal(t, 46, res)
}

func TestV2(t *testing.T) {
	res := solveV2(testCase1)
	require.Equal(t, 51, res)
}
