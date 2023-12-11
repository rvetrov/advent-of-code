package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCaseV1_1 = `
.....
.S-7.
.|.|.
.L-J.
.....
`

const testCaseV1_2 = `
..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`

const testCaseV2_1 = `
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`

const testCaseV2_2 = `
..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........
`

const testCaseV2_3 = `
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`

const testCaseV2_4 = `
FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`

func TestV1(t *testing.T) {
	res := solveV1(testCaseV1_1)
	require.Equal(t, 4, res)

	res = solveV1(testCaseV1_2)
	require.Equal(t, 8, res)
}

func TestV2(t *testing.T) {
	res := solveV2(testCaseV2_1)
	require.Equal(t, 4, res)

	res = solveV2(testCaseV2_2)
	require.Equal(t, 4, res)

	res = solveV2(testCaseV2_3)
	require.Equal(t, 8, res)

	res = solveV2(testCaseV2_4)
	require.Equal(t, 10, res)
}
