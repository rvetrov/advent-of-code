package main

import (
	"fmt"
	"strings"

	"adventofcode.com/2023/internal/utils"
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

func getCode(line string, getDigitFunc func(string) (int, bool)) int {
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

func main() {
	input := utils.MustReadInput("input.big.txt")
	res := 0
	lines := strings.Split(input, "\n")
	fmt.Println("Length", len(lines))
	for _, line := range lines {
		res += getCode(line, getDigitV2)
	}
	utils.MustWriteOutput("output-v2.txt", fmt.Sprint(res))
}
