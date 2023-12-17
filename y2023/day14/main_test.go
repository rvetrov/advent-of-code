package day14

import (
	"testing"

	"github.com/stretchr/testify/require"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

const testCase1 = `
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

const expectedAfterCycles1 = `
.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....

.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O

.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O
`

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 136, res)
}

func TestV2(t *testing.T) {
	field := grid.RotateCCW(utils.NonEmptyLines(testCase1))
	for _, expected := range utils.SplitByEmptyLine(expectedAfterCycles1) {
		field = cycleTilts(field)

		actualField := grid.RotateCW(field)
		require.Equal(t, expected, actualField)
	}

	res := SolveV2(testCase1)
	require.Equal(t, 64, res)
}
