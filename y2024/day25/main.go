package day25

import "adventofcode.com/internal/utils"

const lockRows = 7

func parseLocksAndKeys(input string) (locks [][]int, keys [][]int) {
	for _, block := range utils.EmptyLineSeparatedBlocks(input) {
		if len(block) != lockRows {
			panic(block)
		}
		var pins []int
		for col := range block[0] {
			p := 0
			for row := range block {
				if block[row][col] == '#' {
					p++
				}
			}
			pins = append(pins, p)
		}

		if block[0] == "#####" {
			locks = append(locks, pins)
		} else {
			keys = append(keys, pins)
		}
	}
	return locks, keys
}

func overlaps(lock []int, key []int) bool {
	for i, lockPin := range lock {
		if lockPin+key[i] > lockRows {
			return true
		}
	}
	return false
}

func SolveV1(input string) int {
	locks, keys := parseLocksAndKeys(input)

	res := 0
	for _, lock := range locks {
		for _, key := range keys {
			if !overlaps(lock, key) {
				res++
			}
		}
	}
	return res
}

func SolveV2(input string) int {
	return 0
}
