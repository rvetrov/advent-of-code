package main

import (
	"fmt"

	"adventofcode.com/internal/utils"
)

type Point struct{ X, Y int }

func (p Point) Add(v Vector) Point {
	return Point{p.X + v.DX, p.Y + v.DY}
}

func (p Point) Subtract(other Point) Vector {
	return Vector{p.X - other.X, p.Y - other.Y}
}

type Vector struct{ DX, DY int }

func (v Vector) Reversed() Vector {
	return Vector{-v.DX, -v.DY}
}

const (
	fieldGround = '.'
)

var (
	north = Vector{-1, 0}
	south = Vector{1, 0}
	east  = Vector{0, 1}
	west  = Vector{0, -1}

	directions = map[byte][]Vector{
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
	Start    Point
	grid     []string
	isOnLoop [][]bool
	n        int
	m        int
}

func (f Field) get(p Point) byte {
	if p.X < 0 || f.n <= p.X || p.Y < 0 || f.m <= p.Y {
		return fieldGround
	}
	return f.grid[p.X][p.Y]
}

func (f Field) connected(p1, p2 Point) bool {
	for _, check := range []struct {
		symbol    byte
		direction Vector
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

func (f Field) FindLoop(start Point) int {
	res := 0
	cur := start
	prevDirection := Vector{}
	for {
		f.isOnLoop[cur.X][cur.Y] = true
		prevDirectionReversed := prevDirection.Reversed()
		for _, vec := range []Vector{north, south, west, east} {
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

func (f Field) EnclosedByLoop(start Point) int {
	f.FindLoop(start)
	res := 0
	for i, line := range f.grid {
		insideTheLoop := false

		fmt.Println(line)
		fmt.Println(f.isOnLoop[i])

		for j := 0; j < len(line); j++ {
			if f.isOnLoop[i][j] {
				inVec := Vector{}
				cur := Point{i, j}
				for _, dir := range []Vector{north, south} {
					if f.connected(cur, cur.Add(dir)) {
						inVec = dir
						break
					}
				}

				outVec := inVec.Reversed()
				fmt.Println(j, inVec, outVec)
				if f.connected(cur, cur.Add(outVec)) {
					insideTheLoop = !insideTheLoop
				} else {
					j++
					for {
						cur = Point{i, j}
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
	lines := utils.Lines()
	for _, line := range lines {
		f.grid = append(f.grid, line)
		for j := 0; j < len(line); j++ {
			if line[j] == 'S' {
				f.Start = Point{len(f.grid) - 1, j}
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

func solveV1(input string) int {
	f := parseField(input)
	loopLength := f.FindLoop(f.Start)
	return loopLength / 2
}

func solveV2(input string) int {
	f := parseField(input)
	enclosed := f.EnclosedByLoop(f.Start)
	return enclosed
}

func main() {
	input := utils.MustReadInput("input.big")
	res := solveV2(input)
	utils.MustWriteOutput("output-v2.big", fmt.Sprint(res))
}
