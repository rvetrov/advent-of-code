package day24

import (
	"fmt"
	"strings"

	"adventofcode.com/internal/geom"
	"adventofcode.com/internal/math"
	"adventofcode.com/internal/utils"
)

type Equation []int

func (e Equation) Multiply(x int) {
	for i := range e {
		e[i] *= x
	}
}

func (e Equation) Sub(rhs Equation) {
	if len(e) != len(rhs) {
		panic(fmt.Sprint("Different equation lengths:", e, rhs))
	}
	for i := range e {
		e[i] -= rhs[i]
	}
}

func newEquation(a ...int) Equation {
	return a
}

func intersection2D(p1, p2 geom.PointInt, v1, v2 geom.VectorInt) (geom.PointFloat, float64, float64, bool) {
	// p1X + v1X * t1 = p2X + v2X * t2
	// p1Y * v1Y * t1 = p2Y + v2Y * t2
	// v1X * t1 + (-v2X) * t2 = p2X - p1X
	eqX := newEquation(v1.X, -v2.X, p2.X-p1.X)
	if eqX[0] < 0 {
		eqX.Multiply(-1)
	}
	eqY := newEquation(v1.Y, -v2.Y, p2.Y-p1.Y)
	if eqY[0] < 0 {
		eqY.Multiply(-1)
	}

	// a1 * t1 + a2 * t2 = a3
	// b1 * t1 + b2 * t2 = b3
	lcm := math.LCM(eqX[0], eqY[0])
	if lcm != 0 {
		k1 := lcm / eqX[0]
		k2 := lcm / eqY[0]
		eqX.Multiply(k1)
		eqY.Multiply(k2)
		eqX.Sub(eqY)
	}

	var t1, t2 float64
	if eqX[0] == 0 {
		if eqX[1] == 0 {
			return geom.PointFloat{}, 0, 0, false
		}
		t2 = float64(eqX[2]) / float64(eqX[1])
		t1 = (float64(eqY[2]) - t2*float64(eqY[1])) / float64(eqY[0])
	} else if eqY[0] == 0 {
		if eqY[1] == 0 {
			return geom.PointFloat{}, 0, 0, false
		}
		t2 = float64(eqY[2]) / float64(eqY[1])
		t1 = (float64(eqX[2]) - t2*float64(eqX[1])) / float64(eqX[0])
	} else {
		panic(fmt.Sprint("Error:", eqX, eqY))
	}

	x1 := float64(p1.X) + t1*float64(v1.X)
	y1 := float64(p1.Y) + t1*float64(v1.Y)
	return geom.PointFloat{X: x1, Y: y1}, t1, t2, true
}

func parseInput(input string) (pos []geom.PointInt, vel []geom.VectorInt) {
	for _, line := range utils.NonEmptyLines(input) {
		pStr, vStr, _ := strings.Cut(line, " @ ")
		p := utils.SplitNumbers(pStr, ",")
		pos = append(pos, geom.PointInt{X: p[0], Y: p[1], Z: p[2]})
		v := utils.SplitNumbers(vStr, ",")
		vel = append(vel, geom.VectorInt{X: v[0], Y: v[1], Z: v[2]})
	}
	return pos, vel
}

func solveV1(input string, mn, mx float64) int {
	pos, vel := parseInput(input)

	res := 0
	for i := range len(pos) {
		for j := range i {
			p, t1, t2, found := intersection2D(pos[i], pos[j], vel[i], vel[j])
			if found && t1 >= 0 && t2 >= 0 && mn <= p.X && p.X <= mx && mn <= p.Y && p.Y <= mx {
				res++
			}
		}
	}
	return res
}

func SolveV1(input string) int {
	return solveV1(input, 200000000000000, 400000000000000)
}
