package day01

import "strconv"

func SolveV1(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return res
}

func SolveV2(input string) int {
	return SolveV1(input)
}
