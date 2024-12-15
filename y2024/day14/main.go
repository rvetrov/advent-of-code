package day14

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
	"fmt"
)

type Robot struct {
	pos grid.Position
	vel grid.Direction
}

func parseRobots(lines []string) []Robot {
	var res []Robot
	for _, line := range lines {
		var px, py, vx, vy int
		if cnt, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy); err != nil || cnt != 4 {
			panic(fmt.Sprint(line, cnt, err))
		}
		res = append(res, Robot{
			pos: grid.Position{Row: py, Col: px},
			vel: grid.Direction{DR: vy, DC: vx},
		})
	}
	return res
}

func emulate(robots []Robot, seconds, width, height int) {
	for i := range len(robots) {
		p := robots[i].pos.Add(robots[i].vel.Multiplied(seconds))
		p.Row = (p.Row%height + height) % height
		p.Col = (p.Col%width + width) % width
		robots[i].pos = p
	}
}

func safetyFactor(robots []Robot, width, height int) int {
	midRow := height / 2
	midCol := width / 2
	var cnt [4]int
	for _, r := range robots {
		p := r.pos
		if p.Row == midRow || p.Col == midCol {
			continue
		}
		ind := 0
		if p.Row > midRow {
			ind += 1
		}
		if p.Col > midCol {
			ind += 2
		}
		cnt[ind]++
	}
	//fmt.Println(cnt)
	return cnt[0] * cnt[1] * cnt[2] * cnt[3]
}

func buildRobotMap(robots []Robot, width, height int) [][]byte {
	var gr [][]byte
	for range height {
		row := make([]byte, width)
		for j := range width {
			row[j] = '.'
		}
		gr = append(gr, row)
	}
	for _, r := range robots {
		gr[r.pos.Row][r.pos.Col] = '1'
	}
	return gr
}

func printRobots(robots []Robot, width, height int) {
	gr := buildRobotMap(robots, width, height)
	for _, row := range gr {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func easterEggScore(robots []Robot, width, height int) int {
	gr := buildRobotMap(robots, width, height)
	res := 0
	for i := range len(gr) - 1 {
		for j := range len(gr[0]) - 1 {
			if gr[i][j] != '.' && gr[i+1][j] != '.' {
				res++
			}
			if gr[i][j] != '.' && gr[i][j+1] != '.' {
				res++
			}
		}
	}
	return res
}

func SolveV1(input string) int {
	width, height := 101, 103
	robots := parseRobots(utils.NonEmptyLines(input))
	emulate(robots, 100, width, height)
	return safetyFactor(robots, width, height)
}

func SolveV2(input string) int {
	width, height := 101, 103
	robots := parseRobots(utils.NonEmptyLines(input))
	for sec := range 1000000 {
		emulate(robots, 1, width, height)
		score := easterEggScore(robots, width, height)
		if score > 130 {
			//printRobots(robots, width, height)
			return sec + 1
		}
	}
	return 0
}
