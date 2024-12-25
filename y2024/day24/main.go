package day24

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"adventofcode.com/internal/utils"
)

type Gate struct {
	in1, in2, out int
	op            string
	outStr        string
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

func extractBit(name string) int {
	bit, _ := strconv.Atoi(name[1:])
	return bit
}

type Device struct {
	gates     []Gate
	wireToInd map[string]int
	wireCount int
}

func parseDevice(lines []string) *Device {
	device := &Device{
		wireToInd: make(map[string]int),
	}
	for _, line := range lines {
		g := Gate{}
		parts := strings.Split(line, " ")
		g.in1 = device.WireInd(parts[0])
		g.op = parts[1]
		g.in2 = device.WireInd(parts[2])
		g.out = device.WireInd(parts[4])
		g.outStr = parts[4]
		device.gates = append(device.gates, g)
	}
	return device
}

func (d *Device) Produce(state *big.Int) int64 {
	//gates := d.gates
	//for len(gates) > 0 {
	//	var notProduced []Gate
	//	for _, gate := range gates {
	//		value1, produced1 := state[gate.in1]
	//		value2, produced2 := state[gate.in2]
	//		if !produced1 || !produced2 {
	//			notProduced = append(notProduced, gate)
	//			continue
	//		}
	//		state[gate.out] = gate.produce(value1, value2)
	//	}
	//	gates = notProduced
	//}

	var res int64
	//for name, value := range state {
	//	if name[0] == 'z' && value > 0 {
	//		res |= 1 << extractBit(name)
	//	}
	//}
	return res
}

func (d *Device) dfs(curGate int, deps []map[string]struct{}) map[string]struct{} {
	if deps[curGate] != nil {
		return deps[curGate]
	}
	curDeps := make(map[string]struct{})
	deps[curGate] = curDeps

	//for _, input := range []string{d.gates[curGate].in1, d.gates[curGate].in2} {
	//	if ind, ok := d.wireToInd[input]; ok {
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

func (d *Device) WireInd(wire string) int {
	if ind, ok := d.wireToInd[wire]; ok {
		return ind
	}
	d.wireToInd[wire] = d.wireCount
	d.wireCount++
	return d.wireCount - 1
}

func (d *Device) ParseInput(lines []string) *big.Int {
	var (
		wire  string
		value uint
		res   = big.NewInt(0)
	)
	for _, line := range lines {
		_, _ = fmt.Sscanf(line[1:], "%s %d", &wire, &value)
		wire = wire[:len(wire)-1]
		res.SetBit(res, d.WireInd(wire), value)
	}
	return res
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

func SolveV1(input string) int64 {
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
