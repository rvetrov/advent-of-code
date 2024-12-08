package day08

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

func parseAntennas(input string) (grid.Grid, map[byte][]grid.Position) {
	gr := grid.New(utils.NonEmptyLines(input))
	antennas := make(map[byte][]grid.Position)
	for pos := gr.Start(); gr.Contains(pos); pos = gr.Next(pos) {
		ch, _ := gr.At(pos)
		if ch != '.' {
			antennas[ch] = append(antennas[ch], pos)
		}
	}
	return gr, antennas

}

func SolveV1(input string) int {
	gr, antennas := parseAntennas(input)

	antinodes := make(map[grid.Position]struct{})
	for _, posList := range antennas {
		n := len(posList)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				antinode := posList[i].Add(posList[i].Subtract(posList[j]))
				if gr.Contains(antinode) {
					antinodes[antinode] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}

func SolveV2(input string) int {
	gr, antennas := parseAntennas(input)

	antinodes := make(map[grid.Position]struct{})
	for _, posList := range antennas {
		n := len(posList)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				delta := posList[i].Subtract(posList[j])
				antinode := posList[i]
				for gr.Contains(antinode) {
					antinodes[antinode] = struct{}{}
					antinode = antinode.Add(delta)
				}
			}
		}
	}

	return len(antinodes)
}
