package day01

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestV1(t *testing.T) {
	s := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	res := SolveV1(s)

	require.Equal(t, 142, res)
}

func TestV2(t *testing.T) {
	s := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	res := SolveV2(s)

	require.Equal(t, 281, res)
}
