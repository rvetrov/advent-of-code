package day24

import (
	"fmt"
	"math/bits"
	"slices"
	"strconv"
	"strings"

	"adventofcode.com/internal/utils"
	"golang.org/x/exp/rand"
)

type WireInd int
type GateInd int

const (
	bitUnresolved byte    = 255
	noGate        GateInd = -1
)

type Hint struct {
	name1, name2 string
	ind1, ind2   GateInd
}

type Gate struct {
	in1, in2, out WireInd
	op            string
}

func (g Gate) produce(v1, v2 byte) byte {
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

type InputState [333]byte

func newInputState() InputState {
	var res InputState
	for i := range res {
		res[i] = bitUnresolved
	}
	return res
}

type Device struct {
	gates              []Gate
	wireNameToInd      map[string]WireInd
	wireNames          []string
	wireToBlockedGates [][]GateInd
	wireToGate         []GateInd
	inputBitSize       int
}

func parseDevice(lines []string) *Device {
	d := &Device{
		wireNameToInd: make(map[string]WireInd),
	}
	for _, line := range lines {
		g := Gate{}
		parts := strings.Split(line, " ")
		g.in1 = d.WireInd(parts[0])
		g.op = parts[1]
		g.in2 = d.WireInd(parts[2])
		g.out = d.WireInd(parts[4])

		gateInd := GateInd(len(d.gates))
		d.wireToGate[g.out] = gateInd
		d.wireToBlockedGates[g.in1] = append(d.wireToBlockedGates[g.in1], gateInd)
		d.wireToBlockedGates[g.in2] = append(d.wireToBlockedGates[g.in2], gateInd)
		d.gates = append(d.gates, g)
	}
	return d
}

func (d *Device) WireInd(wire string) WireInd {
	if ind, ok := d.wireNameToInd[wire]; ok {
		return ind
	}
	ind := WireInd(len(d.wireToBlockedGates))
	d.wireNameToInd[wire] = ind
	d.wireToBlockedGates = append(d.wireToBlockedGates, nil)
	d.wireNames = append(d.wireNames, wire)
	d.wireToGate = append(d.wireToGate, noGate)

	if wire[0] == 'x' || wire[0] == 'y' {
		if bit, _ := strconv.Atoi(wire[1:]); d.inputBitSize < bit+1 {
			d.inputBitSize = bit + 1
		}
	}
	return ind
}

func (d *Device) ParseInput(lines []string) InputState {
	var (
		wire  string
		value byte
		res   = newInputState()
	)
	for _, line := range lines {
		_, _ = fmt.Sscanf(line, "%s %d", &wire, &value)
		res[d.WireInd(wire[:len(wire)-1])] = value
	}
	return res
}

func (d *Device) InputNumsToState(x, y int64) InputState {
	var state = newInputState()
	for i := range d.inputBitSize {
		var value byte
		if x&(1<<i) > 0 {
			value = 1
		}
		state[d.WireInd(fmt.Sprintf("x%02d", i))] = value
		value = 0
		if y&(1<<i) > 0 {
			value = 1
		}
		state[d.WireInd(fmt.Sprintf("y%02d", i))] = value
	}
	return state
}

func (d *Device) Produce(state InputState) int64 {
	var ready []GateInd
	produced := make([]bool, len(d.gates))
	for gateInd, gate := range d.gates {
		if state[gate.in1] != bitUnresolved && state[gate.in2] != bitUnresolved {
			ready = append(ready, GateInd(gateInd))
			produced[gateInd] = true
		}
	}

	for len(ready) > 0 {
		gate := d.gates[ready[0]]
		ready = ready[1:]

		value1 := state[gate.in1]
		value2 := state[gate.in2]
		state[gate.out] = gate.produce(value1, value2)

		for _, blocked := range d.wireToBlockedGates[gate.out] {
			if !produced[blocked] && state[d.gates[blocked].in1] != bitUnresolved && state[d.gates[blocked].in2] != bitUnresolved {
				produced[blocked] = true
				ready = append(ready, blocked)
			}
		}
	}

	var res int64
	for wireInd, name := range d.wireNames {
		if name[0] == 'z' && state[wireInd] == 1 {
			bit, _ := strconv.Atoi(name[1:])
			res |= int64(1) << bit
		}
	}
	return res
}

func (d *Device) dfs(swaps int, hints []Hint, prevCorrupted int64, correctProducer func(int64, int64) int64) []GateInd {
	if swaps == 0 {
		if prevCorrupted == 0 {
			return []GateInd{}
		}
		return nil
	}

	for i := range d.gates {
		gateI := GateInd(i)
		var newHints []Hint
		if len(hints) > 0 {
			if hints[0].ind1 != gateI {
				continue
			}
			newHints = hints[1:]
		}

		for j := range len(d.gates) {
			gateJ := GateInd(j)
			if len(hints) > 0 && hints[0].ind2 != gateJ {
				continue
			}

			d.swapGateOuts(gateI, gateJ)
			newCorrupted := detectCorruptedBits(d, 200, correctProducer)

			if prevCorrupted^newCorrupted > 0 && prevCorrupted^newCorrupted == prevCorrupted-newCorrupted {
				fixed := bits.OnesCount64(uint64(prevCorrupted ^ newCorrupted))
				if fixed > 0 {
					if res := d.dfs(swaps-1, newHints, newCorrupted, correctProducer); res != nil {
						d.swapGateOuts(gateI, gateJ)
						res = append(res, gateI, gateJ)
						return res
					}
				}
			}
			d.swapGateOuts(gateI, gateJ)
		}
	}

	return nil
}

func (d *Device) swapGateOuts(i, j GateInd) {
	outI, outJ := d.gates[i].out, d.gates[j].out
	d.gates[i].out, d.gates[j].out = outJ, outI
	d.wireToGate[outI], d.wireToGate[outJ] = d.wireToGate[outJ], d.wireToGate[outI]
}

func (d *Device) explain(gate Gate) string {
	var explainIn1, explainIn2 string
	if gateIn1 := d.wireToGate[gate.in1]; gateIn1 == noGate {
		explainIn1 = fmt.Sprint(d.wireNames[gate.in1])
	} else {
		explainIn1 = d.explain(d.gates[gateIn1])
	}
	if gateIn2 := d.wireToGate[gate.in2]; gateIn2 == noGate {
		explainIn2 = fmt.Sprint(d.wireNames[gate.in2])
	} else {
		explainIn2 = d.explain(d.gates[gateIn2])
	}
	return fmt.Sprintf("(%s %s: %s %s)", d.wireNames[gate.out], gate.op, explainIn1, explainIn2)
}

func detectCorruptedBits(d *Device, iterations int, correctProducer func(int64, int64) int64) int64 {
	var corruptedBits int64
	for range iterations {
		x := rand.Int63() & ((1 << d.inputBitSize) - 1)
		y := rand.Int63() & ((1 << d.inputBitSize) - 1)
		expected := correctProducer(x, y)

		state := d.InputNumsToState(x, y)
		produced := d.Produce(state)

		corruptedBits |= expected ^ produced
	}
	return corruptedBits
}

func explain(d *Device, corruptedBits int64) {
	for _, isResultExpected := range []bool{true, false} {
		var explains []string
		for _, gate := range d.gates {
			outName := d.wireNames[gate.out]

			if outName[0] == 'z' {
				bit, _ := strconv.Atoi(outName[1:])
				if isResultExpected && corruptedBits&(1<<bit) > 0 {
					explains = append(explains, d.explain(gate))
				}
			} else if !isResultExpected && gate.op == "XOR" {
				explains = append(explains, d.explain(gate))
			}
		}

		slices.Sort(explains)
		for _, line := range explains {
			fmt.Println(line)
		}
		fmt.Println()
	}
}

func adjustDevice(d *Device, swaps int, hints []Hint, correctProducer func(int64, int64) int64) string {
	corruptedBits := detectCorruptedBits(d, 500, correctProducer)
	//explain(d, corruptedBits)

	gatesToSwap := d.dfs(swaps, hints, corruptedBits, correctProducer)

	var res []string
	for _, ind := range gatesToSwap {
		res = append(res, d.wireNames[d.gates[ind].out])
	}
	slices.Sort(res)
	return strings.Join(res, ",")
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

	hints := []Hint{
		// found using explain()
		{name1: "z31", name2: "hkh"},
		{name1: "z27", name2: "bfq"},
		{name1: "z18", name2: "hmt"},
		// found using dfs
		{name1: "fjp", name2: "bng"},
	}
	for i := range hints {
		ind1 := device.wireNameToInd[hints[i].name1]
		ind2 := device.wireNameToInd[hints[i].name2]
		hints[i].ind1 = device.wireToGate[ind1]
		hints[i].ind2 = device.wireToGate[ind2]
	}

	return adjustDevice(
		device,
		4,
		hints,
		func(x, y int64) int64 {
			return x + y
		},
	)
}
