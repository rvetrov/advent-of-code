package utils

import (
	"os"
	"strings"
)

func MustReadInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(input)
}

func MustWriteOutput(fileName string, text string) {
	os.WriteFile(fileName, []byte(text), 0644)
}

func Lines(input string) []string {
	res := []string{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			res = append(res, line)
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
