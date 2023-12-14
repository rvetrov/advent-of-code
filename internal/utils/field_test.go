package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testField1 = []string{"qwe", "asd"}
var testField2 = []string{"aq", "sw", "de"}
var testField3 = []string{"dsa", "ewq"}
var testField4 = []string{"ed", "ws", "qa"}

func TestRotation(t *testing.T) {
	field := testField1
	for _, expected := range [][]string{
		testField2,
		testField3,
		testField4,
		testField1,
	} {
		field = RotateCW(field)
		require.Equal(t, expected, field)
	}

	for _, expected := range [][]string{
		testField4,
		testField3,
		testField2,
		testField1,
	} {
		field = RotateCCW(field)
		require.Equal(t, expected, field)
	}
}
