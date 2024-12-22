package day06

import (
	"strings"

	"adventofcode.com/internal/utils"
)

func parseState(input string, join bool) ([]int, []int) {
	lines := utils.NonEmptyLines(input)

	_, line, _ := strings.Cut(lines[0], ":")
	if join {
		line = strings.ReplaceAll(line, " ", "")
	}
	times := utils.SpaceSeparatedIntegers(line)

	_, line, _ = strings.Cut(lines[1], ":")
	if join {
		line = strings.ReplaceAll(line, " ", "")
	}
	distances := utils.SpaceSeparatedIntegers(line)

	return times, distances
}

func waysToWin(time, dist int) int {
	res := 0
	for pressTime := 0; pressTime < time; pressTime++ {
		if pressTime*(time-pressTime) > dist {
			res++
		}
	}
	return res
}

func SolveV1(input string) int {
	times, distances := parseState(input, false)

	res := 1
	for i := range times {
		res *= waysToWin(times[i], distances[i])
	}
	return res
}

func SolveV2(input string) int {
	times, distances := parseState(input, true)

	res := 1
	for i := range times {
		res *= waysToWin(times[i], distances[i])
	}
	return res
}
