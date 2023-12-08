package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"adventofcode.com/internal/utils"
)

type Range struct {
	Start int
	Len   int
}

func (r Range) Contains(x int) bool {
	return r.Start <= x && x < r.Start+r.Len
}

func (r Range) End() int {
	return r.Start + r.Len
}

func (r Range) Intersection(other Range) *Range {
	left := max(r.Start, other.Start)
	right := min(r.End(), other.End())
	if left < right {
		return &Range{
			Start: left,
			Len:   right - left,
		}
	}
	return nil
}

type Ranges []Range

func (rs Ranges) Len() int           { return len(rs) }
func (rs Ranges) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }
func (rs Ranges) Less(i, j int) bool { return rs[i].Start < rs[j].Start }

type MappingRange struct {
	Src Range
	Dst Range
}

func (mp *MappingRange) Contains(x int) bool {
	return mp.Src.Contains(x)
}

func (mp *MappingRange) Map(seed int) int {
	diff := seed - mp.Src.Start
	return mp.Dst.Start + diff
}

type MappingRanges []*MappingRange

func (mp MappingRanges) Len() int {
	return len(mp)
}

func (mp MappingRanges) Swap(i, j int) {
	mp[i], mp[j] = mp[j], mp[i]
}

func (mp MappingRanges) Less(i, j int) bool {
	return mp[i].Src.Start < mp[j].Src.Start
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
			srcStart, dstStart, length := 0, 0, 0
			if scanned, err := fmt.Sscanf(line, "%d %d %d", &dstStart, &srcStart, &length); scanned == 3 && err == nil {
				m := &MappingRange{
					Src: Range{Start: srcStart, Len: length},
					Dst: Range{Start: dstStart, Len: length},
				}
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

func doMapSeeds(seeds Ranges, mapping MappingRanges) Ranges {
	sort.Sort(seeds)
	sort.Sort(mapping)

	newSeeds := Ranges{}
	j := 0
	for i := 0; i < len(seeds); i++ {
		for j < len(mapping) && seeds[i].Len > 0 && mapping[j].Src.Start < seeds[i].Start+seeds[i].Len {
			isec := seeds[i].Intersection(mapping[j].Src)
			// seed:   [___)
			// map : [[[[[[
			if isec == nil {
				// seed:     [___)
				// map : [___)
				j++
			} else if seeds[i].Start < isec.Start {
				// seed: [___)
				// map :   [____
				newSeed := Range{
					Start: seeds[i].Start,
					Len:   isec.Start - seeds[i].Start,
				}
				newSeeds = append(newSeeds, newSeed)
				seeds[i].Len -= newSeed.Len
				seeds[i].Start = isec.Start
			} else {
				// seed:   [___)
				// map : [___
				seeds[i].Start = isec.End()
				seeds[i].Len -= isec.Len
				newSeeds = append(newSeeds, Range{Start: mapping[j].Map(isec.Start), Len: isec.Len})
			}
		}
		if seeds[i].Len > 0 {
			newSeeds = append(newSeeds, seeds[i])
		}
	}

	return newSeeds
}

func solveV2(input string) int {
	seedNums, mappings := parseState(input)
	seeds := Ranges{}
	for i := 0; i < len(seedNums); i += 2 {
		seeds = append(seeds, Range{Start: seedNums[i], Len: seedNums[i+1]})
	}
	fmt.Println(seeds)

	for _, mapRanges := range mappings {
		seeds = doMapSeeds(seeds, mapRanges)
	}
	sort.Sort(seeds)
	return seeds[0].Start
}

func main() {
	input := utils.MustReadInput("input.big")
	res := solveV2(input)
	utils.MustWriteOutput("output-v2.big", fmt.Sprint(res))
}
