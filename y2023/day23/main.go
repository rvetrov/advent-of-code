package day23

import (
	"slices"
	"strings"

	"adventofcode.com/internal/math"
	"adventofcode.com/internal/utils"
)

type Brick struct {
	x1, y1, z1  int
	x2, y2, z2  int
	disappeared bool
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

func parseBricks(input string) []*Brick {
	var res []*Brick
	for _, line := range utils.NonEmptyLines(input) {
		res = append(res, parseBrick(line))
	}
	sortBricks(res)
	return res
}

func sortBricks(bricks []*Brick) {
	slices.SortFunc(bricks, func(b1, b2 *Brick) int {
		if b1.z1 < b2.z1 {
			return -1
		} else if b1.z1 > b2.z1 {
			return 1
		}
		return 0
	})

}

func findFallHeight(a *Brick, bricks []*Brick) int {
	var maxZ = 1
	for _, b := range bricks {
		if b.disappeared {
			continue
		}
		x1 := math.MaxInt(a.x1, b.x1)
		x2 := math.MinInt(a.x2, b.x2)
		if x1 > x2 {
			continue
		}

		y1 := math.MaxInt(a.y1, b.y1)
		y2 := math.MinInt(a.y2, b.y2)
		if y1 > y2 {
			continue
		}

		if maxZ <= b.z2 {
			maxZ = b.z2 + 1
		}
	}
	return maxZ - a.z1
}

func stabilizeBricks(bricks []*Brick) {
	for i, brick := range bricks {
		dz := findFallHeight(brick, bricks[:i])
		brick.z1 += dz
		brick.z2 += dz
	}
	sortBricks(bricks)

}

func SolveV1(input string) int {
	bricks := parseBricks(input)
	stabilizeBricks(bricks)

	res := 0
	for i := range bricks {
		bricks[i].disappeared = true
		safe := true
		for j := i + 1; j < len(bricks); j++ {
			if dz := findFallHeight(bricks[j], bricks[:j]); dz != 0 {
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
	bricks := parseBricks(input)
	stabilizeBricks(bricks)

	res := 0
	for i := range bricks {
		bricks[i].disappeared = true
		disappeared := []int{i}
		for j := i + 1; j < len(bricks); j++ {
			bs := slices.Clone(bricks[:i])
			if dz := findFallHeight(bricks[j], append(bs, bricks[:j]...)); dz != 0 {
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
