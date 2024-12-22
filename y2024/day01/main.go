package day01

import (
	"sort"

	"adventofcode.com/internal/math"
	"adventofcode.com/internal/utils"
)

func readLists(input string) (l1, l2 []int) {
	for _, line := range utils.NonEmptyLines(input) {
		nums := utils.SpaceSeparatedIntegers(line)
		l1 = append(l1, nums[0])
		l2 = append(l2, nums[1])
	}
	return l1, l2
}

func SolveV1(input string) int {
	l1, l2 := readLists(input)

	sort.Ints(l1)
	sort.Ints(l2)

	res := 0
	for i := range len(l1) {
		res += math.AbsInt(l1[i] - l2[i])
	}
	return res
}

func SolveV2(input string) int {
	l1, l2 := readLists(input)

	cnt2 := make(map[int]int)
	for _, num := range l2 {
		cnt2[num] += 1
	}

	res := 0
	for _, num := range l1 {
		res += num * cnt2[num]
	}
	return res
}
