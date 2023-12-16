package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
#.#.### 1,1,3
.#...#....###. 1,1,3
.#.###.#.###### 1,3,1,6
####.#...#... 4,1,1
#....######..#####. 1,6,5
.###.##....# 3,2,1
`

const testCase2 = `
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`

func TestV1(t *testing.T) {
	res := solveV1(testCase1)
	require.Equal(t, 6, res)

	res = solveV1(testCase2)
	require.Equal(t, 21, res)
}

func TestV2(t *testing.T) {
	res := solveV2(testCase1)
	require.Equal(t, 6, res)

	res = solveV2(testCase2)
	require.Equal(t, 525152, res)
}
