package day14

import (
	"strings"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

func tiltLeft(line string) string {
	builder := strings.Builder{}
	j := 0
	bs := strings.Split(line, "")
	for i := 0; i < len(bs); i++ {
		if j <= i {
			j = i + 1
		}
		if bs[i] == "." {
			for j < len(bs) && bs[j] == "." {
				j++
			}
			if j < len(bs) && bs[j] == "O" {
				bs[i], bs[j] = bs[j], bs[i]
			}
		}
		builder.WriteString(bs[i])
	}
	return builder.String()
}

func weighLine(line string) (res int) {
	for i, ch := range line {
		if ch == 'O' {
			res += len(line) - i
		}
	}
	return res
}

func SolveV1(input string) int {
	gr := grid.RotateCCW(utils.NonEmptyLines(input))
	res := 0
	for _, line := range gr {
		res += weighLine(tiltLeft(line))
	}
	return res
}

func cycleTilts(field []string) []string {
	for rotNum := 0; rotNum < 4; rotNum++ {
		for i, line := range field {
			field[i] = tiltLeft(line)
		}
		field = grid.RotateCW(field)
	}
	return field
}

func SolveV2(input string) int {
	gr := grid.RotateCCW(utils.NonEmptyLines(input))
	cache := map[string]int{}

	totalRotations := 1000000000
	for step := 0; step < totalRotations; step++ {
		fingerprint := strings.Join(gr, "")
		if prevStep, ok := cache[fingerprint]; ok {
			cycleLen := step - prevStep
			step += cycleLen * ((totalRotations - step) / cycleLen)
		}
		cache[fingerprint] = step

		gr = cycleTilts(gr)
	}

	res := 0
	for _, line := range gr {
		res += weighLine(line)
	}
	return res
}
