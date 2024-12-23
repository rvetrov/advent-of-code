package day23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testCase1 = `
kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn
`
)

func TestSolveV1(t *testing.T) {
	res := SolveV1(testCase1)
	require.Equal(t, 7, res)
}

func TestSolveV2(t *testing.T) {
	res := SolveV2(testCase1)
	require.Equal(t, "co,de,ka,ta", res)
}
