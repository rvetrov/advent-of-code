package day25

import (
	"fmt"
	"strings"

	"adventofcode.com/internal/utils"
)

type Edge struct {
	From, To string
}

type Graph struct {
	cap  map[string]map[string]int
	flow map[string]map[string]int
}

func NewGraph() Graph {
	return Graph{
		cap:  make(map[string]map[string]int),
		flow: make(map[string]map[string]int),
	}
}

func (g Graph) AddEdge(from, to string) {
	if _, ok := g.cap[from]; !ok {
		g.cap[from] = make(map[string]int)
		g.flow[from] = make(map[string]int)
	}
	g.cap[from][to] = 1
	if _, ok := g.cap[to]; !ok {
		g.cap[to] = make(map[string]int)
		g.flow[to] = make(map[string]int)
	}
	g.cap[to][from] = 1
}

func (g Graph) DeleteEdge(from, to string) {
	delete(g.cap[from], to)
	delete(g.flow[from], to)
	delete(g.cap[to], from)
	delete(g.flow[to], from)
}

func (g Graph) ClearFlow() {
	for vertex := range g.cap {
		if len(g.flow[vertex]) > 0 {
			g.flow[vertex] = make(map[string]int)
		}
	}
}

func parseGraph(input string) (Graph, []Edge) {
	graph := NewGraph()
	var edges []Edge
	for _, line := range utils.NonEmptyLines(input) {
		from, toList, _ := strings.Cut(line, ": ")
		for _, to := range strings.Split(toList, " ") {
			graph.AddEdge(from, to)
			edges = append(edges, Edge{from, to})
		}
	}
	return graph, edges
}

func increaseFlow(graph Graph, from, to string) bool {
	prev := make(map[string]string)
	prev[from] = from
	findFlow(graph, from, prev)
	if _, ok := prev[to]; !ok {
		return false
	}
	for to != from {
		graph.flow[prev[to]][to]++
		graph.flow[to][prev[to]]--
		to = prev[to]
	}
	return true
}

func findFlow(graph Graph, cur string, prev map[string]string) {
	for next, edgeCap := range graph.cap[cur] {
		if _, ok := prev[next]; !ok && graph.flow[cur][next] < edgeCap {
			prev[next] = cur
			findFlow(graph, next, prev)
		}
	}
}

func findComponentSize(current string, graph Graph, visited map[string]struct{}) int {
	visited[current] = struct{}{}
	res := 1
	for adjVertex := range graph.cap[current] {
		if _, v := visited[adjVertex]; !v {
			res += findComponentSize(adjVertex, graph, visited)
		}
	}
	return res
}

func findComponentSizes(graph Graph) []int {
	visited := make(map[string]struct{})
	var res []int
	for vertex := range graph.cap {
		if _, v := visited[vertex]; !v {
			res = append(res, findComponentSize(vertex, graph, visited))
		}
	}
	return res
}

func SolveV1(input string) int {
	graph, edges := parseGraph(input)

	restFlow := 3
	firstEdge := edges[0]
	for _, edge := range edges {
		graph.DeleteEdge(edge.From, edge.To)
		graph.ClearFlow()

		if restFlow == 3 {
			firstEdge = edge
		}

		flow := 0
		for range restFlow {
			if increaseFlow(graph, firstEdge.From, firstEdge.To) {
				flow++
			}
		}
		if flow == restFlow-1 {
			fmt.Println("Found:", edge)
			restFlow--
		} else {
			graph.AddEdge(edge.From, edge.To)
		}
	}

	sizes := findComponentSizes(graph)
	fmt.Println(sizes)
	if len(sizes) == 2 {
		return sizes[0] * sizes[1]
	}
	panic("not found")
}
