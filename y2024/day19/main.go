package day19

import (
	"strings"

	"adventofcode.com/internal/utils"
)

func possibleDesigns(design string, towels []string) int {
	designs := make([]int, len(design)+1)
	designs[0] = 1
	for i := range len(design) + 1 {
		if designs[i] == 0 {
			continue
		}
		for _, towel := range towels {
			if i+len(towel) <= len(design) && towel == design[i:i+len(towel)] {
				designs[i+len(towel)] += designs[i]
			}
		}
	}
	return designs[len(design)]
}

func SolveV1(input string) int {
	blocks := utils.EmptyLineSeparatedBlocks(input)
	towels := strings.Split(blocks[0][0], ", ")

	res := 0
	for _, design := range blocks[1] {
		if possibleDesigns(design, towels) > 0 {
			res += 1
		}
	}
	return res
}

func SolveV2(input string) int {
	blocks := utils.EmptyLineSeparatedBlocks(input)
	towels := strings.Split(blocks[0][0], ", ")

	res := 0
	for _, design := range blocks[1] {
		res += possibleDesigns(design, towels)
	}
	return res
}
