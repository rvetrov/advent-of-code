package day03

import (
	"fmt"
	"regexp"
	"strconv"

	"adventofcode.com/internal/utils"
)

type Instruction string

const (
	doInst   Instruction = "do()"
	dontInst Instruction = "don't()"
)

type expression struct {
	tokensRegExp *regexp.Regexp
	mulRegExp    *regexp.Regexp
	enabled      bool
}

func newExpression() *expression {
	return &expression{enabled: true}
}

func (e *expression) extractInstruction(line string) []Instruction {
	if e.tokensRegExp == nil {
		e.tokensRegExp = regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)
	}
	var res []Instruction
	for _, token := range e.tokensRegExp.FindAllString(line, -1) {
		res = append(res, Instruction(token))
	}
	return res
}

func (e *expression) evalMul(inst Instruction) int {
	if e.mulRegExp == nil {
		e.mulRegExp = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	}
	groups := e.mulRegExp.FindStringSubmatch(string(inst))
	if len(groups) != 3 {
		panic(groups)
	}
	n1, err1 := strconv.Atoi(groups[1])
	n2, err2 := strconv.Atoi(groups[2])
	if err1 != nil || err2 != nil {
		panic(fmt.Sprint(inst, err1, err2))
	}
	return n1 * n2
}

func (e *expression) EvalMuls(line string, enabling bool) int {
	var res int
	for _, inst := range e.extractInstruction(line) {
		switch inst {
		case doInst:
			if enabling {
				e.enabled = true
			}
		case dontInst:
			if enabling {
				e.enabled = false
			}
		default:
			if e.enabled {
				res += e.evalMul(inst)
			}
		}
	}
	return res
}

func SolveV1(input string) int {
	var res int
	expr := newExpression()
	for _, line := range utils.NonEmptyLines(input) {
		res += expr.EvalMuls(line, false)
	}
	return res
}

func SolveV2(input string) int {
	var res int
	expr := newExpression()
	for _, line := range utils.NonEmptyLines(input) {
		res += expr.EvalMuls(line, true)
	}
	return res
}
