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

func MustReadInput(fileName string) string {
	input, err := ReadInput(fileName)
	if err != nil {
		panic(err)
	}
	return input
}

func MustWriteOutput(fileName string, text string) {
	os.WriteFile(fileName, []byte(text), 0644)
}

func NonEmptyLines(input string) []string {
	res := []string{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			res = append(res, line)
		}
	}
	return res
}

func ReadNumbers(line string) []int {
	res := []int{}
	for _, numStr := range strings.Split(line, " ") {
		if num, err := strconv.Atoi(numStr); err == nil {
			res = append(res, num)
		}
	}
	return res
}

func SplitByEmptyLine(input string) [][]string {
	res := [][]string{}

	field := []string{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			field = append(field, line)
		} else {
			if len(field) > 0 {
				res = append(res, field)
			}
			field = []string{}
		}
	}
	if len(field) > 0 {
		res = append(res, field)
	}
	return res
}
