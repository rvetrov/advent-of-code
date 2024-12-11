package day11

import (
	"fmt"
	"strconv"

	"adventofcode.com/internal/utils"
)

func readStones(input string) map[int]int {
	res := make(map[int]int)
	for _, stone := range utils.ReadNumbers(utils.NonEmptyLines(input)[0]) {
		res[stone]++
	}
	return res
}

func countStones(stones map[int]int) int {
	res := 0
	for _, value := range stones {
		res += value
	}
	return res
}

func blink(stones map[int]int) map[int]int {
	res := make(map[int]int)
	for stone, cnt := range stones {
		if stone == 0 {
			res[1] += cnt
			continue
		}
		numStr := fmt.Sprint(stone)
		if len(numStr)%2 == 0 {
			half := len(numStr) / 2
			left, _ := strconv.Atoi(numStr[:half])
			right, _ := strconv.Atoi(numStr[half:])
			res[left] += cnt
			res[right] += cnt
		} else {
			res[stone*2024] += cnt
		}
	}
	return res
}

func solve(input string, blinks int) int {
	stones := readStones(input)
	for range blinks {
		stones = blink(stones)
	}
	return countStones(stones)
}

func SolveV1(input string) int {
	return solve(input, 25)
}

func SolveV2(input string) int {
	return solve(input, 75)
}
