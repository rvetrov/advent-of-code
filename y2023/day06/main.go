package day06

import (
	"fmt"
	"strings"

	"adventofcode.com/internal/utils"
)

func parseState(input string, join bool) ([]int, []int) {
	lines := utils.NonEmptyLines(input)

	_, line, _ := strings.Cut(lines[0], ":")
	fmt.Println(lines[0])
	fmt.Println(line)
	if join {
		line = strings.ReplaceAll(line, " ", "")
	}
	times := utils.ReadNumbers(line)
	fmt.Println(line)

	_, line, _ = strings.Cut(lines[1], ":")
	fmt.Println(lines[1])
	fmt.Println(line)
	if join {
		line = strings.ReplaceAll(line, " ", "")
	}
	distances := utils.ReadNumbers(line)
	fmt.Println(line)

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

func solveV1(input string) int {
	times, distances := parseState(input, false)

	res := 1
	for i := range times {
		res *= waysToWin(times[i], distances[i])
	}
	return res
}

func solveV2(input string) int {
	times, distances := parseState(input, true)

	res := 1
	for i := range times {
		res *= waysToWin(times[i], distances[i])
	}
	return res
}
