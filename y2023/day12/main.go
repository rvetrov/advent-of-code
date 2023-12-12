package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"adventofcode.com/internal/utils"
)

func parseField(line string) (string, []int) {
	field, groupsStr, _ := strings.Cut(line, " ")
	groups := []int{}
	for _, groupStr := range strings.Split(groupsStr, ",") {
		if groupLen, err := strconv.Atoi(groupStr); err == nil {
			groups = append(groups, groupLen)
		} else {
			log.Fatalf("Failed to parse %q: %q", groupStr, err)
		}
	}
	return field, groups
}

func canBePlaced(length, pos int, field string) bool {
	if pos+length > len(field) {
		return false
	}
	for i := 0; i < length; i++ {
		if field[pos+i] == '.' {
			return false
		}
	}
	return pos+length == len(field) || field[pos+length] != '#'
}

func numberOfPossibleArrangements(field string, groups []int) int {
	n := len(field)
	resPrev := make([]int, n+1)
	resPrev[0] = 1
	for i := 0; i < n; i++ {
		if field[i] != '#' {
			resPrev[i+1] = resPrev[i]
		} else {
			break
		}
	}

	for _, group := range groups {
		res := make([]int, n+1)
		for i := 0; i < n; i++ {
			if field[i] != '#' {
				res[i+1] += res[i]
			}

			if resPrev[i] > 0 && canBePlaced(group, i, field) {
				j := i + group + 1
				if j > n {
					j = n
				}
				res[j] += resPrev[i]
			}
		}

		resPrev = res
	}

	return resPrev[n]
}

func solveV1(input string) int {
	lines := utils.Lines(input)
	res := 0
	for _, line := range lines {
		field, groups := parseField(line)
		res += numberOfPossibleArrangements(field, groups)
	}
	return res
}

func solveV2(input string) int {
	lines := utils.Lines(input)
	res := 0
	for _, line := range lines {
		field, groups := parseField(line)
		fs := []string{}
		gs := []int{}
		for i := 0; i < 5; i++ {
			fs = append(fs, field)
			gs = append(gs, groups...)
		}
		field = strings.Join(fs, "?")
		res += numberOfPossibleArrangements(field, gs)
	}
	return res
}

func main() {
	input := utils.MustReadInput("input.big")
	res := solveV2(input)
	utils.MustWriteOutput("output-v2.big", fmt.Sprint(res))
}
