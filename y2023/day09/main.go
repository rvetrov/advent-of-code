package day09

import (
	"log"
	"strconv"
	"strings"
)

type History []int

func (h History) allZeroes() bool {
	for _, num := range h {
		if num != 0 {
			return false
		}
	}
	return true
}

func (h History) predictNext() int {
	if h.allZeroes() {
		return 0
	}
	nextH := make(History, len(h)-1)
	for j := 0; j+1 < len(h); j++ {
		nextH[j] = h[j+1] - h[j]
	}
	return nextH.predictNext() + h[len(h)-1]
}

func (h History) predictPrev() int {
	if h.allZeroes() {
		return 0
	}
	nextH := make(History, len(h)-1)
	for j := 0; j+1 < len(h); j++ {
		nextH[j] = h[j+1] - h[j]
	}
	return h[0] - nextH.predictPrev()
}

func parseHistories(input string) []History {
	res := []History{}
	for _, line := range strings.Split(input, "\n") {
		if line = strings.TrimSpace(line); len(line) > 0 {
			history := History{}
			for _, numStr := range strings.Split(line, " ") {
				if num, err := strconv.Atoi(numStr); err == nil {
					history = append(history, num)
				} else {
					log.Fatalf("Failed to parse %q: %q", numStr, err)
				}
			}
			res = append(res, history)
		}
	}
	return res
}

func solveV1(input string) int {
	hs := parseHistories(input)
	res := 0
	for _, history := range hs {
		res += history.predictNext()
	}
	return res
}

func solveV2(input string) int {
	hs := parseHistories(input)
	res := 0
	for _, history := range hs {
		res += history.predictPrev()
	}
	return res
}
