package day25

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase1 = `
jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr
`
)

func TestV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 54, res)
}
