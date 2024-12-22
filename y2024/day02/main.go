package day02

import (
	"slices"

	"adventofcode.com/internal/utils"
)

func isSafeIncreasingPair(a, b int) bool {
	return a < b && b <= a+3
}

func isSafeIncreasingReport(nums []int, allowBadLevel bool) bool {
	if !allowBadLevel {
		for i := 1; i < len(nums); i++ {
			if !isSafeIncreasingPair(nums[i-1], nums[i]) {
				return false
			}
		}
		return true
	}

	if isSafeIncreasingReport(nums[1:], false) {
		return true
	}
	if isSafeIncreasingReport(nums[:len(nums)-1], false) {
		return true
	}

	for i := 1; i+1 < len(nums); i++ {
		if isSafeIncreasingPair(nums[i-1], nums[i+1]) &&
			(!isSafeIncreasingPair(nums[i-1], nums[i]) || !isSafeIncreasingPair(nums[i], nums[i+1])) &&
			isSafeIncreasingReport(nums[:i], false) &&
			isSafeIncreasingReport(nums[i+1:], false) {
			return true
		}
	}
	return false
}

func isSafeReport(nums []int, allowBadLevel bool) bool {
	if len(nums) < 2 {
		return true
	}
	if isSafeIncreasingReport(nums, allowBadLevel) {
		return true
	}
	slices.Reverse(nums)
	return isSafeIncreasingReport(nums, allowBadLevel)
}

func SolveV1(input string) int {
	res := 0
	for _, line := range utils.NonEmptyLines(input) {
		nums := utils.SpaceSeparatedIntegers(line)
		if isSafeReport(nums, false) {
			res += 1
		}
	}
	return res
}

func SolveV2(input string) int {
	res := 0
	for _, line := range utils.NonEmptyLines(input) {
		nums := utils.SpaceSeparatedIntegers(line)
		if isSafeReport(nums, true) {
			res += 1
		}
	}
	return res
}
