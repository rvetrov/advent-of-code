package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testCase1 = `
rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
`

func TestV1(t *testing.T) {
	res := solveV1(testCase1)
	require.Equal(t, 1320, res)
}

func TestV2(t *testing.T) {
	res := solveV2(testCase1)
	require.Equal(t, 145, res)
}
