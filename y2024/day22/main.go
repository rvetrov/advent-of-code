package day22

import (
	"strconv"

	"adventofcode.com/internal/utils"
)

const pruningModule = 16777216

func nextRandom(prev int) int {
	prev = ((prev * 64) ^ prev) % pruningModule
	prev = ((prev / 32) ^ prev) % pruningModule
	return ((prev * 2048) ^ prev) % pruningModule
}

type fourChanges struct {
	c1, c2, c3, c4 int
}

func simulateRandom(rnd, iterations int) int {
	for range iterations {
		rnd = nextRandom(rnd)
	}
	return rnd
}

func updateChangeGains(rnd, iterations int, changeGain map[fourChanges]int) {
	changeExists := make(map[fourChanges]struct{})
	last4 := fourChanges{}
	prevPrice := rnd % 10
	for it := range iterations {
		rnd = nextRandom(rnd)
		price := rnd % 10

		last4.c1, last4.c2, last4.c3, last4.c4 = last4.c2, last4.c3, last4.c4, price-prevPrice
		if it >= 3 {
			if _, exists := changeExists[last4]; !exists {
				changeExists[last4] = struct{}{}
				changeGain[last4] += price
			}
		}
		prevPrice = price
	}
}

func SolveV1(input string) int {
	res := 0
	for _, numStr := range utils.NonEmptyLines(input) {
		num, _ := strconv.Atoi(numStr)
		res += simulateRandom(num, 2000)
	}
	return res
}

func SolveV2(input string) int {
	changeGains := make(map[fourChanges]int)
	for _, numStr := range utils.NonEmptyLines(input) {
		num, _ := strconv.Atoi(numStr)
		updateChangeGains(num, 2000, changeGains)
	}

	res := 0
	for _, gain := range changeGains {
		if res < gain {
			res = gain
		}
	}
	return res
}
