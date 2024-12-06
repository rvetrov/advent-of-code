package grid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testGrid1 = New([]string{"qwe", "asd"})
var testGrid2 = New([]string{"aq", "sw", "de"})
var testGrid3 = New([]string{"dsa", "ewq"})
var testGrid4 = New([]string{"ed", "ws", "qa"})

func TestRotation(t *testing.T) {
	gr := testGrid1
	for _, expected := range []Grid{
		testGrid2,
		testGrid3,
		testGrid4,
		testGrid1,
	} {
		gr = RotateCW(gr)
		require.Equal(t, expected, gr)
	}

	for _, expected := range []Grid{
		testGrid4,
		testGrid3,
		testGrid2,
		testGrid1,
	} {
		gr = RotateCCW(gr)
		require.Equal(t, expected, gr)
	}
}
