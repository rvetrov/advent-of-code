package day10

import (
	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

const (
	fieldGround = '.'
)

var (
	north = grid.Up
	south = grid.Down
	east  = grid.Right
	west  = grid.Left

	directions = map[byte][]grid.Direction{
		'|':         {north, south},
		'-':         {west, east},
		'L':         {north, east},
		'J':         {north, west},
		'7':         {south, west},
		'F':         {south, east},
		'S':         {south, north, west, east},
		fieldGround: {},
	}
)

type Field struct {
	Start    grid.Position
	grid     []string
	isOnLoop [][]bool
	n        int
	m        int
}

func (f Field) get(p grid.Position) byte {
	if p.Row < 0 || f.n <= p.Row || p.Col < 0 || f.m <= p.Col {
		return fieldGround
	}
	return f.grid[p.Row][p.Col]
}

func (f Field) connected(p1, p2 grid.Position) bool {
	for _, check := range []struct {
		symbol    byte
		direction grid.Direction
	}{
		{f.get(p1), p2.Subtract(p1)},
		{f.get(p2), p1.Subtract(p2)},
	} {
		found := false
		for _, direction := range directions[check.symbol] {
			if check.direction == direction {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (f Field) FindLoop(start grid.Position) int {
	res := 0
	cur := start
	prevDirection := grid.Direction{}
	for {
		f.isOnLoop[cur.Row][cur.Col] = true
		prevDirectionReversed := prevDirection.Reversed()
		for _, vec := range []grid.Direction{north, south, west, east} {
			cand := cur.Add(vec)
			if vec != prevDirectionReversed && f.connected(cur, cand) {
				cur = cand
				prevDirection = vec
				break
			}
		}
		res++
		if cur == start {
			break
		}
	}
	return res
}

func (f Field) EnclosedByLoop(start grid.Position) int {
	f.FindLoop(start)
	res := 0
	for i, line := range f.grid {
		insideTheLoop := false

		for j := 0; j < len(line); j++ {
			if f.isOnLoop[i][j] {
				inVec := grid.Direction{}
				cur := grid.Position{Row: i, Col: j}
				for _, dir := range []grid.Direction{north, south} {
					if f.connected(cur, cur.Add(dir)) {
						inVec = dir
						break
					}
				}

				outVec := inVec.Reversed()
				if f.connected(cur, cur.Add(outVec)) {
					insideTheLoop = !insideTheLoop
				} else {
					j++
					for {
						cur = grid.Position{Row: i, Col: j}
						if f.connected(cur, cur.Add(outVec)) {
							insideTheLoop = !insideTheLoop
							break
						} else if f.connected(cur, cur.Add(inVec)) {
							break
						} else {
							j++
						}
					}
				}

			} else if insideTheLoop {
				res++
			}
		}
	}
	return res
}

func parseField(input string) Field {
	f := Field{grid: []string{}}
	lines := utils.NonEmptyLines(input)
	for _, line := range lines {
		f.grid = append(f.grid, line)
		for j := 0; j < len(line); j++ {
			if line[j] == 'S' {
				f.Start = grid.Position{Row: len(f.grid) - 1, Col: j}
			}
		}
	}
	f.n = len(f.grid)
	f.m = len(f.grid[0])
	for i := 0; i < f.n; i++ {
		f.isOnLoop = append(f.isOnLoop, make([]bool, f.m))
	}
	return f
}

func SolveV1(input string) int {
	f := parseField(input)
	loopLength := f.FindLoop(f.Start)
	return loopLength / 2
}

func SolveV2(input string) int {
	f := parseField(input)
	enclosed := f.EnclosedByLoop(f.Start)
	return enclosed
}
