package day23

import (
	"slices"
	"sort"
	"strings"

	"adventofcode.com/internal/utils"
)

type Network struct {
	links   [][]bool
	nodeInd map[string]int
	nodes   []string
}

func parseNetwork(edges []string) *Network {
	n := &Network{
		nodeInd: make(map[string]int),
	}
	for _, edge := range edges {
		p1, p2, _ := strings.Cut(edge, "-")

		if _, ok := n.nodeInd[p1]; !ok {
			n.nodeInd[p1] = len(n.nodes)
			n.nodes = append(n.nodes, p1)
		}
		if _, ok := n.nodeInd[p2]; !ok {
			n.nodeInd[p2] = len(n.nodes)
			n.nodes = append(n.nodes, p2)
		}
	}

	for range n.nodes {
		n.links = append(n.links, make([]bool, len(n.nodes)))
	}

	for _, edge := range edges {
		p1, p2, _ := strings.Cut(edge, "-")
		p1Ind := n.nodeInd[p1]
		p2Ind := n.nodeInd[p2]
		n.links[p1Ind][p2Ind] = true
		n.links[p2Ind][p1Ind] = true
	}
	return n
}

func SolveV1(input string) int {
	n := parseNetwork(utils.NonEmptyLines(input))

	res := 0
	for i1, name1 := range n.nodes {
		for i2, name2 := range n.nodes[:i1] {
			if !n.links[i1][i2] {
				continue
			}
			for i3, name3 := range n.nodes[:i2] {
				if n.links[i3][i1] && n.links[i3][i2] && (name1[0] == 't' || name2[0] == 't' || name3[0] == 't') {
					res++
				}
			}
		}
	}
	return res
}

func findClique(n *Network, nodes []int, start int) []int {
	res := nodes
	for ; start < len(n.nodes); start++ {
		ok := true
		for _, node := range nodes {
			if !n.links[start][node] {
				ok = false
				break
			}
		}
		if ok {
			clique := findClique(n, append(nodes, start), start+1)
			if len(res) < len(clique) {
				res = slices.Clone(clique)
			}
		}
	}
	return res
}

func SolveV2(input string) string {
	n := parseNetwork(utils.NonEmptyLines(input))
	clique := findClique(n, nil, 0)

	var names []string
	for _, ind := range clique {
		names = append(names, n.nodes[ind])
	}
	sort.Strings(names)
	return strings.Join(names, ",")
}
