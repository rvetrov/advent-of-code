package day14

import (
	"github.com/stretchr/testify/require"
	"testing"

	"adventofcode.com/internal/utils"
)

const (
	testCase0 = `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`
)

func TestSolveV1(t *testing.T) {
	width, height := 11, 7
	robots := parseRobots(utils.NonEmptyLines(testCase0))
	//printRobots(robots, width, height)
	emulate(robots, 100, width, height)
	//printRobots(robots, width, height)
	sf := safetyFactor(robots, width, height)
	require.Equal(t, 12, sf)
}

func TestSolveV2(t *testing.T) {
}
