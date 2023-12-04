package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"adventofcode.com/2023/internal/utils"
)

func parseNumbers(str string) []int {
	res := []int{}
	for _, numStr := range strings.Split(str, " ") {
		if num, err := strconv.Atoi(numStr); err == nil {
			res = append(res, num)
		}
	}
	return res
}

func cardScore(winHaveStr string) int {
	winningStr, youHaveStr, _ := strings.Cut(winHaveStr, "|")
	winningNums := parseNumbers(winningStr)
	youHaveNums := parseNumbers(youHaveStr)

	res := 0
	for _, num := range youHaveNums {
		if slices.Contains(winningNums, num) {
			res++
		}
	}
	return res
}

func splitAllCards(input string) []string {
	res := []string{}
	for _, cardStr := range strings.Split(input, "\n") {
		if _, winHaveStr, found := strings.Cut(cardStr, ":"); found {
			res = append(res, winHaveStr)
		}
	}
	return res
}

func solveV1(input string) int {
	res := 0
	for _, cardStr := range splitAllCards(input) {
		score := cardScore(cardStr)
		if score > 0 {
			res += 1 << (score - 1)
		}
	}
	return res
}

func solveV2(input string) int {
	cards := splitAllCards(input)
	var scores = make([]int, len(cards))
	for i := range scores {
		scores[i] = 1
	}

	res := 0
	for i, cardStr := range cards {
		res += scores[i]
		times := cardScore(cardStr)
		if times > 0 {
			for j := 1; j <= times; j++ {
				scores[i+j] += scores[i]
			}
		}
	}
	return res
}

func main() {
	input := utils.MustReadInput("input.big")
	res := solveV2(input)
	utils.MustWriteOutput("output-v2.big", fmt.Sprint(res))
}
