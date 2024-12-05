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

func ReadNumbers(line string) []int {
	var res []int
	for _, numStr := range strings.Split(line, " ") {
		if num, err := strconv.Atoi(numStr); err == nil {
			res = append(res, num)
		}
	}
	return res
}

func SplitByEmptyLine(input string) [][]string {
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
