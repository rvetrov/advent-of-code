package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLowerBound(t *testing.T) {
	for n := 1; n <= 10; n++ {
		var stack []int
		for i := range n {
			stack = append(stack, i)
		}

		for needle := -1; needle <= n+1; needle++ {
			foundPos := LowerBound(
				0,
				n,
				func(ind int) bool {
					return stack[ind] >= needle
				},
			)
			expected := needle
			if expected < 0 {
				expected = 0
			} else if expected > n {
				expected = n
			}

			require.Equal(t, expected, foundPos)
		}
	}
}
