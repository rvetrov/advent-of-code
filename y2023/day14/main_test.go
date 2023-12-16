package day14

import (
	"testing"

	"adventofcode.com/internal/utils"
	"github.com/stretchr/testify/require"
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
	res := solveV1(testCase1)
	require.Equal(t, 136, res)
}

func TestV2(t *testing.T) {
	field := utils.RotateCCW(utils.NonEmptyLines(testCase1))
	for _, expected := range utils.SplitByEmptyLine(expectedAfterCycles1) {
		field = cycleTilts(field)

		actualField := utils.RotateCW(field)
		require.Equal(t, expected, actualField)
	}

	res := solveV2(testCase1)
	require.Equal(t, 64, res)
}
