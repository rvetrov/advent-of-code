package day20

import (
	"strings"

	"adventofcode.com/internal/utils"
)

type System struct {
	ms      map[string]*Module
	HighCnt int
	LowCnt  int

	SandMachineTurnedOn bool
}

func NewSystem(lines []string) *System {
	s := System{ms: map[string]*Module{}}
	for _, line := range lines {
		m := NewModule(line)
		s.ms[m.Name] = m
	}
	for _, m := range s.ms {
		for _, dstName := range m.Dsts {
			if dst, found := s.ms[dstName]; found {
				dst.Srcs = append(dst.Srcs, m.Name)
			}
		}
	}
	return &s
}

func (s *System) PushTheButton() {
	q := []Pulse{{Src: "button", Dst: string(mtBroadcaster), Type: pulseLow}}
	for len(q) > 0 {
		pulse := q[0]
		if pulse.Type == pulseLow && pulse.Dst == "rx" {
			s.SandMachineTurnedOn = true
		}
		q = q[1:]
		if pulse.Type == pulseHigh {
			s.HighCnt++
		} else {
			s.LowCnt++
		}

		if m, found := s.ms[pulse.Dst]; found {
			q = append(q, m.Receive(pulse.Type, pulse.Src)...)
		}
	}
}

type PulseType int
type ModuleType string

const (
	mtBroadcaster ModuleType = "broadcaster"
	mtFlipFlop    ModuleType = "%"
	mtConjunction ModuleType = "&"

	pulseLow  PulseType = 0
	pulseHigh PulseType = 1
)

type Module struct {
	Name  string
	Type  ModuleType
	Dsts  []string
	Srcs  []string
	State int

	highPulseSrcs map[string]bool
	highPulses    int
}

func NewModule(line string) *Module {
	m := Module{
		highPulseSrcs: map[string]bool{},
	}
	before, after, _ := strings.Cut(line, " -> ")

	if before == string(mtBroadcaster) {
		m.Type = mtBroadcaster
		m.Name = string(mtBroadcaster)
	} else {
		m.Type = ModuleType(before[:1])
		m.Name = before[1:]
	}

	m.Dsts = strings.Split(after, ", ")
	return &m
}

func (m *Module) Receive(pulse PulseType, from string) []Pulse {
	res := []Pulse{}
	if m.Type == mtFlipFlop {
		if pulse == pulseLow {
			m.State = 1 - m.State
			pulseToSend := pulseLow
			if m.State == 1 {
				pulseToSend = pulseHigh
			}
			for _, dst := range m.Dsts {
				res = append(res, Pulse{Type: pulseToSend, Src: m.Name, Dst: dst})
			}
		}
	} else if m.Type == mtConjunction {
		prevIsHigh := m.highPulseSrcs[from]
		newIsHigh := pulse == pulseHigh
		m.highPulseSrcs[from] = newIsHigh
		if newIsHigh && !prevIsHigh {
			m.highPulses++
		} else if !newIsHigh && prevIsHigh {
			m.highPulses--
		}

		pulseToSend := pulseHigh
		if len(m.Srcs) == m.highPulses {
			pulseToSend = pulseLow
		}
		for _, dst := range m.Dsts {
			res = append(res, Pulse{Type: pulseToSend, Src: m.Name, Dst: dst})
		}
	} else if m.Type == mtBroadcaster {
		for _, dst := range m.Dsts {
			res = append(res, Pulse{Type: pulse, Src: m.Name, Dst: dst})
		}
	} else {
		panic(m)
	}
	return res
}

type Pulse struct {
	Type PulseType
	Src  string
	Dst  string
}

func solve(input string, pushes int) int {
	lines := utils.NonEmptyLines(input)
	s := NewSystem(lines)

	for i := 0; i < pushes; i++ {
		s.PushTheButton()
	}
	return s.HighCnt * s.LowCnt
}

func SolveV1(input string) int {
	return solve(input, 1000)
}

func SolveV2(input string) int {
	lines := utils.NonEmptyLines(input)
	s := NewSystem(lines)

	res := 0
	for !s.SandMachineTurnedOn && res < 100000 {
		res++
		s.PushTheButton()
	}
	return res
}
