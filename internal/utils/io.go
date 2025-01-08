package utils

import (
	"os"
	"strconv"
	"strings"
)

func ReadInput(fileName string) (string, error) {
	if input, err := os.ReadFile(fileName); err != nil {
		return "", err
	} else {
		return string(input), nil
	}
}

func NonEmptyLines(input string) []string {
	var res []string
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			res = append(res, line)
		}
	}
	return res
}

func SplitNumbers(line, sep string) []int {
	var res []int
	for _, numStr := range strings.Split(line, sep) {
		numStr = strings.TrimSpace(numStr)
		if num, err := strconv.Atoi(numStr); err == nil {
			res = append(res, num)
		}
	}
	return res
}

func SpaceSeparatedIntegers(line string) []int {
	return SplitNumbers(line, " ")
}

func EmptyLineSeparatedBlocks(input string) [][]string {
	var res [][]string

	var textBlock []string
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			textBlock = append(textBlock, line)
		} else {
			if len(textBlock) > 0 {
				res = append(res, textBlock)
			}
			textBlock = []string{}
		}
	}
	if len(textBlock) > 0 {
		res = append(res, textBlock)
	}
	return res
}
