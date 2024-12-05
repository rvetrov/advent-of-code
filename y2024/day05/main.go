package day05

import (
	"slices"
	"strconv"
	"strings"

	"adventofcode.com/internal/utils"
)

type orderRules struct {
	before, after map[string]map[string]struct{}
}

func parseOrderRules(ruleLines []string) *orderRules {
	rules := &orderRules{
		before: make(map[string]map[string]struct{}),
		after:  make(map[string]map[string]struct{}),
	}
	for _, rule := range ruleLines {
		ruleParts := strings.Split(rule, "|")
		p1, p2 := ruleParts[0], ruleParts[1]

		if _, ok := rules.before[p2]; !ok {
			rules.before[p2] = make(map[string]struct{})
		}
		rules.before[p2][p1] = struct{}{}
		if _, ok := rules.after[p1]; !ok {
			rules.after[p1] = make(map[string]struct{})
		}
		rules.after[p1][p2] = struct{}{}
	}
	return rules
}

func (or *orderRules) correctOrder(update []string) bool {
	for i, u1 := range update {
		for _, u2 := range update[i+1:] {
			if _, found := or.before[u1][u2]; found {
				return false
			}
		}
	}
	return true
}

func (or *orderRules) reorder(update []string) []string {
	slices.SortFunc(update, func(p1, p2 string) int {
		if _, ok := or.before[p2][p1]; ok {
			return -1
		}
		if _, ok := or.before[p1][p2]; ok {
			return 1
		}
		return 0
	})
	return update
}

func parseUpdates(lines []string) [][]string {
	var updates [][]string
	for _, line := range lines {
		updates = append(updates, strings.Split(line, ","))
	}
	return updates
}

func SolveV1(input string) int {
	blocks := utils.SplitByEmptyLine(input)
	rules := parseOrderRules(blocks[0])

	res := 0
	for _, update := range parseUpdates(blocks[1]) {
		if rules.correctOrder(update) {
			num, _ := strconv.Atoi(update[len(update)/2])
			res += num
		}
	}
	return res
}

func SolveV2(input string) int {
	blocks := utils.SplitByEmptyLine(input)
	rules := parseOrderRules(blocks[0])

	res := 0
	for _, update := range parseUpdates(blocks[1]) {
		if !rules.correctOrder(update) {
			update = rules.reorder(update)
			num, _ := strconv.Atoi(update[len(update)/2])
			res += num
		}
	}
	return res
}
