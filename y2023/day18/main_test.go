package day18

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
`

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 62, res)
}

func TestV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, 0, res)
	// require.Equal(t, 952408144115, res)
}
