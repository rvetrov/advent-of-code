package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"adventofcode.com/2023/internal/utils"
)

type MappingRange struct {
	SrcStart int
	DstStart int
	Length   int
}

func (mp *MappingRange) Contains(seed int) bool {
	return mp.SrcStart <= seed && seed < mp.SrcStart+mp.Length
}

func (mp *MappingRange) Map(seed int) int {
	diff := seed - mp.SrcStart
	return mp.DstStart + diff
}

type MappingRanges []*MappingRange

func (mp MappingRanges) Len() int {
	return len(mp)
}

func (mp MappingRanges) Swap(i, j int) {
	mp[i], mp[j] = mp[j], mp[i]
}

func (mp MappingRanges) Less(i, j int) bool {
	return mp[i].SrcStart < mp[j].SrcStart
}

func parseState(input string) ([]int, []MappingRanges) {
	lines := strings.Split(input, "\n")
	seeds := []int{}
	mappings := make([]MappingRanges, 0)
	lastInd := -1
	for _, line := range lines {

		if before, after, found := strings.Cut(line, ":"); found {
			if before == "seeds" {
				for _, seedStr := range strings.Split(after, " ") {
					if seed, err := strconv.Atoi(seedStr); err == nil {
						seeds = append(seeds, seed)
					}
				}
			} else {
				mappings = append(mappings, MappingRanges{})
				lastInd++
			}
		} else {
			m := &MappingRange{}
			if scanned, err := fmt.Sscanf(line, "%d %d %d", &m.DstStart, &m.SrcStart, &m.Length); scanned == 3 && err == nil {
				mappings[lastInd] = append(mappings[lastInd], m)
			}
		}
	}
	return seeds, mappings

}

func solveV1(input string) int {
	seeds, mappings := parseState(input)
	fmt.Println(seeds)
	fmt.Println(mappings)

	for _, mapRanges := range mappings {
		newSeeds := make([]int, len(seeds))
		for i, seed := range seeds {
			newSeeds[i] = seed
			for _, mapRange := range mapRanges {
				if mapRange.Contains(seed) {
					newSeeds[i] = mapRange.Map(seed)
					break
				}
			}
		}
		seeds = newSeeds
	}
	res := seeds[0]
	for _, seed := range seeds {
		if res > seed {
			res = seed
		}
	}
	return res
}

func doMapSeeds(seeds MappingRanges, mapping MappingRanges) MappingRanges {
	sort.Sort(seeds)
	sort.Sort(mapping)

	newSeeds := MappingRanges{}
	j := 0
	for i := 0; i < len(seeds); i++ {
		for j < len(mapping) && mapping[j].SrcStart < seeds[i].SrcStart+seeds[i].Length && seeds[i].Length > 0 {
			// seed:   [___)
			// map : [[[[[[
			if mapping[j].SrcStart+mapping[j].Length <= seeds[i].SrcStart {
				// seed:     [___)
				// map : [___)
				j++
			} else if seeds[i].SrcStart < mapping[j].SrcStart {
				// seed: [___)
				// map :   [____
				newSeed := &MappingRange{
					SrcStart: seeds[i].SrcStart,
					Length:   mapping[j].SrcStart - seeds[i].SrcStart,
				}
				newSeeds = append(newSeeds, newSeed)
				seeds[i].Length -= newSeed.Length
				seeds[i].SrcStart = mapping[j].SrcStart
			} else {
				// seed:   [___)
				// map : [___
				newStart := seeds[i].SrcStart
				newFin := newStart + seeds[i].Length
				if newFin > mapping[j].SrcStart+mapping[j].Length {
					newFin = mapping[j].SrcStart + mapping[j].Length
				}
				seeds[i].SrcStart = newFin
				seeds[i].Length -= newFin - newStart
				newSeeds = append(newSeeds, &MappingRange{
					SrcStart: mapping[j].Map(newStart),
					Length:   newFin - newStart,
				})
			}
		}
		if seeds[i].Length > 0 {
			newSeeds = append(newSeeds, seeds[i])
		}
	}

	return newSeeds
}

func solveV2(input string) int {
	seedNums, mappings := parseState(input)
	seeds := MappingRanges{}
	for i := 0; i < len(seedNums); i += 2 {
		seeds = append(seeds, &MappingRange{SrcStart: seedNums[i], Length: seedNums[i+1]})
	}
	fmt.Println(seeds)
	fmt.Println(mappings)

	for _, mapRanges := range mappings {
		seeds = doMapSeeds(seeds, mapRanges)
	}
	sort.Sort(seeds)
	return seeds[0].SrcStart
}

func main() {
	input := utils.MustReadInput("input.big")
	res := solveV2(input)
	utils.MustWriteOutput("output-v2.big", fmt.Sprint(res))
}
