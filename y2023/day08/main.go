package main

import (
	"fmt"

	"adventofcode.com/internal/math"
	"adventofcode.com/internal/utils"
)

type Crossroad struct {
	Left  string
	Right string
}

type Map struct {
	Route string
	Edges map[string]Crossroad
}

type MileStone struct {
	Pos  string
	Step int
}

type Route []MileStone

func (m Map) FindLoop(position string) (int, Route) {
	visited := map[string]int{}

	totalRouteSteps := 0
	stateStr := fmt.Sprintf("%s,%d", position, totalRouteSteps)
	visited[stateStr] = totalRouteSteps
	route := Route{{stateStr, totalRouteSteps}}
	loopSize := 0

	for {
		routeStep := totalRouteSteps % len(m.Route)
		if m.Route[routeStep] == 'L' {
			position = m.Edges[position].Left
		} else {
			position = m.Edges[position].Right
		}
		totalRouteSteps++

		stateStr = fmt.Sprintf("%s,%d", position, totalRouteSteps%len(m.Route))
		if stateStr[2] == 'Z' {
			route = append(route, MileStone{stateStr, totalRouteSteps})
			if visited[stateStr] > 0 {
				loopSize = totalRouteSteps - visited[stateStr]
				break
			}
			visited[stateStr] = totalRouteSteps
		}
	}
	return loopSize, route
}

func newMap(input string) Map {
	lines := utils.Lines(input)

	edges := map[string]Crossroad{}
	for i := 1; i < len(lines); i++ {
		var from, left, right string
		if n, err := fmt.Sscanf(lines[i], "%3s = (%3s, %3s)", &from, &left, &right); err == nil {
			edges[from] = Crossroad{left, right}
		} else {
			fmt.Println(n, err)
		}
	}
	return Map{Route: lines[0], Edges: edges}
}

func solveV1(input string) int {
	mp := newMap(input)
	res := 0
	for pos, cur := 0, "AAA"; cur != "ZZZ"; pos = (pos + 1) % len(mp.Route) {
		if mp.Route[pos] == 'L' {
			cur = mp.Edges[cur].Left
		} else {
			cur = mp.Edges[cur].Right
		}
		res++
	}
	return res
}

func solveV2(input string) int {
	mp := newMap(input)
	ghosts := []string{}
	for from := range mp.Edges {
		if from[2] == 'A' {
			ghosts = append(ghosts, from)
		}
	}
	fmt.Println("Ghosts:", len(ghosts))

	res := 1
	for _, ghost := range ghosts {
		loopSize, route := mp.FindLoop(ghost)
		fmt.Println(ghost, loopSize, route)
		// Prints the following for input.big:
		// XCA 19099 [{XCA,0 0} {NNZ,0 19099} {NNZ,0 38198}]
		// AAA 19637 [{AAA,0 0} {ZZZ,0 19637} {ZZZ,0 39274}]
		// QGA 14257 [{QGA,0 0} {GHZ,0 14257} {GHZ,0 28514}]
		// LBA 11567 [{LBA,0 0} {NPZ,0 11567} {NPZ,0 23134}]
		// GSA 12643 [{GSA,0 0} {SPZ,0 12643} {SPZ,0 25286}]
		// LHA 15871 [{LHA,0 0} {HVZ,0 15871} {HVZ,0 31742}]

		// Every route is a loop with only one target position. Therefore, LCM works.
		res = math.LCM(res, loopSize)
	}
	return res
}

func main() {
	input := utils.MustReadInput("input.big")
	res := solveV2(input)
	utils.MustWriteOutput("output-v2.big", fmt.Sprint(res))
}
