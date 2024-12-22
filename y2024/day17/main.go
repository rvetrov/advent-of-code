package day17

import (
	"fmt"
	"slices"
	"strings"

	"adventofcode.com/internal/utils"
)

type Computer struct {
	steps int
	inst  int
	regA  int
	regB  int
	regC  int
	out   []int
}

func (c *Computer) combo(literal int) int {
	if literal == 4 {
		return c.regA
	} else if literal == 5 {
		return c.regB
	} else if literal == 6 {
		return c.regC
	}
	return literal
}

func (c *Computer) op(opcode, literal int) bool {

	divOp := func(a, b int) int {
		for a > 0 && b > 0 {
			a /= 2
			b--
		}
		return a
	}

	c.steps++
	combo := c.combo(literal)
	switch opcode {
	case 0:
		c.regA = divOp(c.regA, combo)
	case 1:
		c.regB ^= literal
	case 2:
		c.regB = combo % 8
	case 3:
		if c.regA != 0 {
			c.inst = literal
			return false
		}
	case 4:
		c.regB ^= c.regC
	case 5:
		c.out = append(c.out, combo%8)
	case 6:
		c.regB = divOp(c.regA, combo)
	case 7:
		c.regC = divOp(c.regA, combo)
	}
	return true
}

func (c *Computer) Compute(program []int) {
	for c.inst = 0; c.inst+1 < len(program); {
		if c.op(program[c.inst], program[c.inst+1]) {
			c.inst += 2
		}
	}
}

func (c *Computer) Out() string {
	var res []string
	for _, out := range c.out {
		res = append(res, fmt.Sprint(out))
	}
	return strings.Join(res, ",")
}

func NewComputer(lines []string) *Computer {
	c := &Computer{}
	_, _ = fmt.Sscanf(lines[0], "Register A: %d", &c.regA)
	_, _ = fmt.Sscanf(lines[1], "Register B: %d", &c.regB)
	_, _ = fmt.Sscanf(lines[2], "Register C: %d", &c.regC)
	return c
}

func SolveV1(input string) string {
	blocks := utils.EmptyLineSeparatedBlocks(input)
	comp := NewComputer(blocks[0])
	_, programStr, _ := strings.Cut(blocks[1][0], ": ")
	program := utils.SplitNumbers(programStr, ",")

	comp.Compute(program)
	return comp.Out()
}

func digitsToInt(digits []int) int {
	res := 0
	for i, digit := range digits {
		res += digit << (i * 3)
	}
	return res
}

func dfs(digits []int, program []int, compInitLines []string) []int {
	comp := NewComputer(compInitLines)
	comp.regA = digitsToInt(digits)
	comp.Compute(program)

	if slices.Equal(program, comp.out) {
		return digits
	}
	if len(program) < len(comp.out) {
		return nil
	}

	lenToCheck := len(digits) - 2
	for digit := range 8 {
		candDigits := append(digits, digit)

		candInt := digitsToInt(candDigits)
		comp = NewComputer(compInitLines)
		comp.regA = candInt
		comp.Compute(program)

		//fmt.Println("Candidate:", candDigits, candInt)
		//fmt.Println(comp.out, program[:lenToCheck])
		//fmt.Println()

		if len(program) >= lenToCheck && len(comp.out) >= lenToCheck && slices.Equal(program[:lenToCheck], comp.out[:lenToCheck]) {
			if res := dfs(candDigits, program, compInitLines); res != nil {
				return res
			}
		}
	}
	return nil
}

func SolveV2(input string) int {
	blocks := utils.EmptyLineSeparatedBlocks(input)
	_, programStr, _ := strings.Cut(blocks[1][0], ": ")
	program := utils.SplitNumbers(programStr, ",")

	res := -1
	for firstDigit := range 8 {
		for secondDigit := range 8 {
			for thirdDigit := range 8 {
				resDigits := dfs([]int{firstDigit, secondDigit, thirdDigit}, program, blocks[0])
				if len(resDigits) > 0 {
					resCand := digitsToInt(resDigits)
					//fmt.Println("Found:", resDigits, resCand)
					if res == -1 || resCand < res {
						res = resCand
					}
				}
			}
		}
	}
	return res
}
