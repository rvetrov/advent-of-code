package day07

import (
	"strconv"
	"strings"

	"adventofcode.com/internal/utils"
)

func parseEquation(line string) (int, []int) {
	parts := strings.Split(line, ":")
	result, _ := strconv.Atoi(parts[0])
	return result, utils.ReadNumbers(parts[1])
}

func canBeProducedByTwoOps(result int, nums []int) bool {
	if len(nums) == 1 {
		return result == nums[0]
	}
	last := nums[len(nums)-1]
	if result%last == 0 && canBeProducedByTwoOps(result/last, nums[:len(nums)-1]) {
		return true
	}
	return canBeProducedByTwoOps(result-last, nums[:len(nums)-1])
}

func canBeProducedByThreeOps(result int, nums []int) bool {
	if len(nums) == 1 {
		return result == nums[0]
	}
	last := nums[len(nums)-1]
	if result%last == 0 && canBeProducedByThreeOps(result/last, nums[:len(nums)-1]) {
		return true
	}
	if canBeProducedByThreeOps(result-last, nums[:len(nums)-1]) {
		return true
	}

	for last > 0 && result > 0 {
		if last%10 != result%10 {
			return false
		}
		last /= 10
		result /= 10
	}
	return canBeProducedByThreeOps(result, nums[:len(nums)-1])
}

func SolveV1(input string) int {
	res := 0
	for _, line := range utils.NonEmptyLines(input) {
		eqRes, nums := parseEquation(line)
		if canBeProducedByTwoOps(eqRes, nums) {
			res += eqRes
		}
	}
	return res
}

func SolveV2(input string) int {
	res := 0
	for _, line := range utils.NonEmptyLines(input) {
		eqRes, nums := parseEquation(line)
		if canBeProducedByThreeOps(eqRes, nums) {
			res += eqRes
		}
	}
	return res
}
