package day15

import (
	"log"
	"strconv"
	"strings"

	"adventofcode.com/internal/utils"
)

func stringHash(s string) int {
	res := 0
	for _, ch := range s {
		res = ((res + int(ch)) * 17) % 256
	}
	return res
}

const (
	opTypeRemove = '-'
	opTypeSet    = '='
)

type Box struct {
	slots []*Lens
	index map[string]int
}

func (b *Box) Apply(op Operation) {
	if b.index == nil {
		b.index = make(map[string]int)
	}
	if op.Type == opTypeRemove {
		if i, ok := b.index[op.Lens.Label]; ok {
			b.slots[i] = nil
			delete(b.index, op.Lens.Label)
		}
	} else {
		if i, ok := b.index[op.Lens.Label]; ok {
			b.slots[i] = &op.Lens
		} else {
			b.index[op.Lens.Label] = len(b.slots)
			b.slots = append(b.slots, &op.Lens)
		}
	}
}

func (b *Box) FocusingPower() int {
	res := 0
	realInd := 0
	for _, lens := range b.slots {
		if lens != nil {
			realInd++
			res += realInd * lens.FocalLen
		}
	}
	return res
}

type Lens struct {
	Label    string
	FocalLen int
}

type Operation struct {
	Raw  string
	Type rune
	Lens Lens
}

func NewOperation(s string) Operation {
	op := Operation{Raw: s}
	if s[len(s)-1] == opTypeRemove {
		op.Lens.Label = s[:len(s)-1]
		op.Type = opTypeRemove
	} else {
		op.Type = opTypeSet
		label, numStr, _ := strings.Cut(s, string(opTypeSet))
		op.Lens.Label = label
		if num, err := strconv.Atoi(numStr); err == nil {
			op.Lens.FocalLen = num
		} else {
			log.Fatalf("Failed to parse %v", err)
		}
	}
	return op
}

func parseOperations(input string) []Operation {
	ops := []Operation{}
	for _, line := range utils.NonEmptyLines(input) {
		for _, s := range strings.Split(line, ",") {
			ops = append(ops, NewOperation(s))
		}
	}
	return ops
}

func SolveV1(input string) int {
	res := 0
	for _, op := range parseOperations(input) {
		res += stringHash(op.Raw)
	}
	return res
}

func SolveV2(input string) int {
	boxes := make([]Box, 256)
	for _, op := range parseOperations(input) {
		boxI := stringHash(op.Lens.Label)
		boxes[boxI].Apply(op)
	}

	res := 0
	for i, box := range boxes {
		res += box.FocusingPower() * (i + 1)
	}
	return res
}
