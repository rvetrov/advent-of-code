package day01

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestV1(t *testing.T) {
	s := `1`

	res := SolveV1(s)

	require.Equal(t, 1, res)
}

func TestV2(t *testing.T) {
	s := `2`

	res := SolveV2(s)

	require.Equal(t, 2, res)
}
