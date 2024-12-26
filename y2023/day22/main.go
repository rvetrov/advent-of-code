package day22

import (
	"slices"
	"strings"

	"adventofcode.com/internal/utils"
)

type Brick struct {
	x1, y1, z1  int
	x2, y2, z2  int
	disappeared bool
	legs        []int
}

func parseBrick(line string) *Brick {
	b1, b2, _ := strings.Cut(line, "~")
	b1Nums := utils.SplitNumbers(b1, ",")
	b2Nums := utils.SplitNumbers(b2, ",")
	if b1Nums[0] > b2Nums[0] {
		b1Nums[0], b2Nums[0] = b2Nums[0], b1Nums[0]
	}
	if b1Nums[1] > b2Nums[1] {
		b1Nums[1], b2Nums[1] = b2Nums[1], b1Nums[1]
	}
	if b1Nums[2] > b2Nums[2] {
		b1Nums[2], b2Nums[2] = b2Nums[2], b1Nums[2]
	}

	return &Brick{
		x1: b1Nums[0],
		y1: b1Nums[1],
		z1: b1Nums[2],
		x2: b2Nums[0],
		y2: b2Nums[1],
		z2: b2Nums[2],
	}
}

func parseAndSortBricks(input string) []*Brick {
	var res []*Brick
	for _, line := range utils.NonEmptyLines(input) {
		res = append(res, parseBrick(line))
	}
	sortBricks(res)
	return res
}

func sortBricks(bricks []*Brick) {
	slices.SortFunc(bricks, func(a, b *Brick) int {
		if a.z1 < b.z1 {
			return -1
		} else if a.z1 > b.z1 {
			return 1
		}
		return 0
	})
}

func findFallHeight(a *Brick, bricks []*Brick) (int, []int) {
	var maxZ = 1
	var legs []int
	for bi, b := range bricks {
		if b.disappeared {
			continue
		}
		if a.x2 < b.x1 || b.x2 < a.x1 || a.y2 < b.y1 || b.y2 < a.y1 {
			continue
		}

		if maxZ == b.z2+1 {
			legs = append(legs, bi)
		} else if maxZ < b.z2+1 {
			maxZ = b.z2 + 1
			legs = []int{bi}
		}
	}
	return maxZ - a.z1, legs
}

func stabilizeBricks(bricks []*Brick) {
	for i, brick := range bricks {
		dz, _ := findFallHeight(brick, bricks[:i])
		brick.z1 += dz
		brick.z2 += dz
	}
	sortBricks(bricks)
	for i, brick := range bricks {
		_, legs := findFallHeight(brick, bricks[:i])
		brick.legs = legs
	}
}

func SolveV1(input string) int {
	bricks := parseAndSortBricks(input)
	stabilizeBricks(bricks)

	res := 0
	for i := range bricks {
		safe := true
		for j := i + 1; j < len(bricks); j++ {
			if len(bricks[j].legs) == 1 && bricks[j].legs[0] == i {
				safe = false
				break
			}
		}
		if safe {
			res++
		}
		bricks[i].disappeared = false
	}
	return res
}

func SolveV2(input string) int {
	bricks := parseAndSortBricks(input)
	stabilizeBricks(bricks)

	res := 0
	for i := range bricks {
		bricks[i].disappeared = true
		disappeared := []int{i}
		for j := i + 1; j < len(bricks); j++ {
			var legsDisappeared = len(bricks[j].legs) > 0
			for _, leg := range bricks[j].legs {
				if !bricks[leg].disappeared {
					legsDisappeared = false
					break
				}
			}
			if legsDisappeared {
				res++
				bricks[j].disappeared = true
				disappeared = append(disappeared, j)
			}
		}
		for _, ind := range disappeared {
			bricks[ind].disappeared = false
		}
	}
	return res
}
