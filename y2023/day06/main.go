package main

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode.com/internal/utils"
)

func parseState(input string) ([]int, []int) {
	times := []int{}
	distances := []int{}
	lines := []string{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	for _, parse := range []struct {
		To   *[]int
		Line string
	}{
		{&times, lines[0]},
		{&distances, lines[1]},
	} {
		line, _ := strings.CutPrefix(parse.Line, ":")
		for _, numStr := range strings.Split(line, " ") {
			if num, err := strconv.Atoi(numStr); err == nil {
				*parse.To = append(*parse.To, num)
			}
		}
	}

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
	times, distances := parseState(input)
	fmt.Println(times)
	fmt.Println(distances)

	res := 1
	for i := range times {
		res *= waysToWin(times[i], distances[i])
	}
	return res
}

func solveV2(input string) int {
	return solveV1(input)
}

func main() {
	input := utils.MustReadInput("input-v2.big")
	res := solveV1(input)
	utils.MustWriteOutput("output-v2.big", fmt.Sprint(res))
}
