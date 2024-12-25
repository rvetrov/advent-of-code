package day24

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode.com/internal/utils"
)

type Gate struct {
	in1, in2, out int
	op            string
}

func (g Gate) produce(v1, v2 int) int {
	switch g.op {
	case "AND":
		return v1 & v2
	case "OR":
		return v1 | v2
	case "XOR":
		return v1 ^ v2
	default:
		panic(g)
	}
}

type InputState map[int]int

type Device struct {
	gates         []Gate
	wireNameToInd map[string]int
	wireNames     []string
	wireGates     [][]int
}

func parseDevice(lines []string) *Device {
	d := &Device{
		wireNameToInd: make(map[string]int),
	}
	for _, line := range lines {
		g := Gate{}
		parts := strings.Split(line, " ")
		g.in1 = d.WireInd(parts[0])
		g.op = parts[1]
		g.in2 = d.WireInd(parts[2])
		g.out = d.WireInd(parts[4])

		gateInd := len(d.gates)
		d.wireGates[g.in1] = append(d.wireGates[g.in1], gateInd)
		d.wireGates[g.in2] = append(d.wireGates[g.in2], gateInd)
		d.gates = append(d.gates, g)
	}
	return d
}

func (d *Device) WireInd(wire string) int {
	if ind, ok := d.wireNameToInd[wire]; ok {
		return ind
	}
	ind := len(d.wireGates)
	d.wireNameToInd[wire] = ind
	d.wireGates = append(d.wireGates, nil)
	d.wireNames = append(d.wireNames, wire)
	return ind
}

func (d *Device) ParseInput(lines []string) InputState {
	var (
		wire  string
		value int
		res   = make(InputState)
	)
	for _, line := range lines {
		_, _ = fmt.Sscanf(line, "%s %d", &wire, &value)
		res[d.WireInd(wire[:len(wire)-1])] = value
	}
	return res
}

func (d *Device) Produce(state InputState) int {
	//big.NewInt(0)
	//res.SetBit(res, d.WireInd(wire), value)

	var ready []int
	for i, gate := range d.gates {
		_, produced1 := state[gate.in1]
		_, produced2 := state[gate.in2]
		if produced1 && produced2 {
			ready = append(ready, i)
		}
	}

	handled := 0
	for len(ready) > 0 {
		gate := d.gates[ready[0]]
		ready = ready[1:]
		handled++

		value1, _ := state[gate.in1]
		value2, _ := state[gate.in2]
		state[gate.out] = gate.produce(value1, value2)

		for _, blocked := range d.wireGates[gate.out] {
			_, produced1 := state[d.gates[blocked].in1]
			_, produced2 := state[d.gates[blocked].in2]
			if produced1 && produced2 {
				ready = append(ready, blocked)
			}
		}
	}
	if handled != len(d.gates) {
		panic("not all gates handled")
	}

	var res int
	for wireInd, value := range state {
		if d.wireNames[wireInd][0] == 'z' && value > 0 {
			bit, _ := strconv.Atoi(d.wireNames[wireInd][1:])
			res |= 1 << bit
		}
	}
	return res
}

func (d *Device) dfs(curGate int, deps []map[string]struct{}) map[string]struct{} {
	if deps[curGate] != nil {
		return deps[curGate]
	}
	curDeps := make(map[string]struct{})
	deps[curGate] = curDeps

	//for _, input := range []string{d.gates[curGate].in1, d.gates[curGate].in2} {
	//	if ind, ok := d.wireNameToInd[input]; ok {
	//		inputDeps := d.dfs(ind, deps)
	//		for in := range inputDeps {
	//			curDeps[in] = struct{}{}
	//		}
	//	} else {
	//		curDeps[input] = struct{}{}
	//	}
	//}

	return curDeps
}

func (d *Device) analyze() {
	deps := make([]map[string]struct{}, len(d.gates))
	for i := range d.gates {
		if deps[i] == nil {
			d.dfs(i, deps)
		}
		//if d.gates[i].out[0] == 'z' {
		//	fmt.Println(d.gates[i], len(deps[i]))
		//}
	}
}

func adjustDevice(d *Device, swaps int, correctProducer func(int, int) int) string {
	d.analyze()

	corruptedBits := make(map[int]struct{})
	for range 10 {
		//a1

	}
	fmt.Println(corruptedBits)

	return ""
}

func SolveV1(input string) int {
	blocks := utils.EmptyLineSeparatedBlocks(input)
	device := parseDevice(blocks[1])
	state := device.ParseInput(blocks[0])
	return device.Produce(state)
}

func SolveV2(input string) string {
	blocks := utils.EmptyLineSeparatedBlocks(input)
	device := parseDevice(blocks[1])
	return adjustDevice(
		device,
		4,
		func(x, y int) int {
			return x + y
		},
	)
}
