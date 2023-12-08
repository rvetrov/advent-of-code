package main

import (
	"fmt"
	"strings"

	"adventofcode.com/internal/utils"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getDigitV1(line string) (int, bool) {
	if '0' <= line[0] && line[0] <= '9' {
		return int(line[0] - '0'), true
	}
	return 0, false
}

func getDigitV2(line string) (int, bool) {
	if digit, ok := getDigitV1(line); ok {
		return digit, true
	}
	for digit, digitStr := range digits {
		if strings.HasPrefix(line, digitStr) {
			return digit + 1, true
		}
	}
	return 0, false
}

type getDigitFuncDecl func(string) (int, bool)

func getCode(line string, getDigitFunc getDigitFuncDecl) int {
	found := false
	first := 0
	last := 0
	for i := range line {
		if digit, ok := getDigitFunc(line[i:]); ok {
			if !found {
				found = true
				first = digit
			}
			last = digit
		}
	}
	return first*10 + last
}

func solve(input string, getDigitFunc getDigitFuncDecl) int {
	res := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		res += getCode(line, getDigitFunc)
	}
	return res
}

func main() {
	input := utils.MustReadInput("input.big.txt")
	res := solve(input, getDigitV2)
	utils.MustWriteOutput("output-v2.txt", fmt.Sprint(res))
}
