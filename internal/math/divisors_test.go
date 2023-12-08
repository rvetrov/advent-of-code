package math

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGCD(t *testing.T) {
	for _, tc := range []struct{ a, b, gcd int }{
		{8, 12, 4},
		{54, 24, 6},
		{48, 18, 6},
	} {
		require.Equal(t, tc.gcd, GCD(tc.a, tc.b))
	}
}

func TestLCM(t *testing.T) {
	for _, tc := range []struct{ a, b, lcm int }{
		{1, 17, 17},
		{21, 6, 42},
	} {
		require.Equal(t, tc.lcm, LCM(tc.a, tc.b))
	}
}
