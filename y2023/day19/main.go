package day19

import (
	"fmt"
	"log"
	"maps"
	"strings"

	"adventofcode.com/internal/utils"
)

type Operation string

const (
	opLess = "<"
	opMore = ">"
	opNo   = "_"

	opResultAccepted = "A"
	opResultRejected = "R"
)

type Rule struct {
	Op       Operation
	Category string
	Value    int
	Result   string
}

func NewRule(s string) Rule {
	// Example: a<2006:qkq,m>2090:A,rfg
	r := Rule{}
	if before, after, found := strings.Cut(s, ":"); found {
		r.Result = after
		if scanned, err := fmt.Sscanf(before, "%1s%1s%d", &r.Category, &r.Op, &r.Value); scanned != 3 || err != nil {
			log.Fatalln(before, scanned, err)
		}
	} else {
		r.Op = opNo
		r.Result = before
	}
	return r
}

func (r Rule) Apply(p Part) string {
	if r.Op == opNo {
		return r.Result
	}
	target := p[r.Category]
	if r.Op == opLess && target < r.Value || r.Op == opMore && target > r.Value {
		return r.Result
	}
	return ""
}

type Part map[string]int

func NewPart(line string) Part {
	// Example: {x=787,m=2655,a=1222,s=2876}
	var x, m, a, s int
	if n, err := fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s); n != 4 || err != nil {
		log.Fatalln(line, n, err)
	}
	return Part{"x": x, "m": m, "a": a, "s": s}
}

func (p Part) TotalRating() int {
	res := 0
	for _, rating := range p {
		res += rating
	}
	return res
}

type PartRange struct {
	Min map[string]int
	Max map[string]int
}

type Workflow struct {
	Name  string
	Rules []Rule
}

func NewWorkflow(line string) Workflow {
	// Example: hfm{a<1891:R,m<2881:A,m<2987:A,R}

	wf := Workflow{}
	if name, rules, found := strings.Cut(line, "{"); !found {
		log.Fatalln(line)
	} else {
		wf.Name = name
		for _, ruleStr := range strings.Split(rules[:len(rules)-1], ",") {
			wf.Rules = append(wf.Rules, NewRule(ruleStr))
		}
	}
	return wf
}

func (wf Workflow) Next(p Part) string {
	for _, rule := range wf.Rules {
		if applyRes := rule.Apply(p); applyRes != "" {
			return applyRes
		}
	}
	return ""
}

type Workflows map[string]Workflow

func NewWorkflows(lines []string) Workflows {
	wfs := Workflows{}
	for _, line := range lines {
		wf := NewWorkflow(line)
		wfs[wf.Name] = wf
	}
	return wfs
}

func (wfs Workflows) Accepts(p Part) bool {
	cur := "in"
	for cur != opResultAccepted && cur != opResultRejected {
		cur = wfs[cur].Next(p)
	}
	return cur == opResultAccepted
}

func (wfs Workflows) LookupRange(wf string, ruleInd int, pr PartRange) int {
	if wf == opResultAccepted {
		return (pr.Max["x"] - pr.Min["x"] + 1) * (pr.Max["m"] - pr.Min["m"] + 1) *
			(pr.Max["a"] - pr.Min["a"] + 1) * (pr.Max["s"] - pr.Min["s"] + 1)
	} else if wf == opResultRejected {
		return 0
	}

	rule := wfs[wf].Rules[ruleInd]
	if rule.Op == opNo {
		return wfs.LookupRange(rule.Result, 0, pr)
	}

	mn := pr.Min[rule.Category]
	mx := pr.Max[rule.Category]

	res := 0
	if rule.Op == opLess {
		if mn < rule.Value {
			newPR := PartRange{
				Min: maps.Clone(pr.Min),
				Max: maps.Clone(pr.Max),
			}
			newPR.Max[rule.Category] = min(rule.Value-1, mx)
			res += wfs.LookupRange(rule.Result, 0, newPR)
		}
		if rule.Value <= mx {
			newPR := PartRange{
				Min: maps.Clone(pr.Min),
				Max: maps.Clone(pr.Max),
			}
			newPR.Min[rule.Category] = max(rule.Value, mn)
			res += wfs.LookupRange(wf, ruleInd+1, newPR)
		}
	} else if rule.Op == opMore {
		if rule.Value < mx {
			newPR := PartRange{
				Min: maps.Clone(pr.Min),
				Max: maps.Clone(pr.Max),
			}
			newPR.Min[rule.Category] = max(rule.Value+1, mn)
			res += wfs.LookupRange(rule.Result, 0, newPR)
		}
		if mn <= rule.Value {
			newPR := PartRange{
				Min: maps.Clone(pr.Min),
				Max: maps.Clone(pr.Max),
			}
			newPR.Max[rule.Category] = min(rule.Value, mx)
			res += wfs.LookupRange(wf, ruleInd+1, newPR)
		}
	} else {
		panic(rule)
	}

	return res
}

func SolveV1(input string) int {
	blocks := utils.EmptyLineSeparatedBlocks(input)
	wfs := NewWorkflows(blocks[0])

	res := 0
	for _, line := range blocks[1] {
		part := NewPart(line)
		if wfs.Accepts(part) {
			res += part.TotalRating()
		}
	}
	return res
}

func SolveV2(input string) int {
	blocks := utils.EmptyLineSeparatedBlocks(input)
	wfs := NewWorkflows(blocks[0])

	start := PartRange{
		Min: map[string]int{"x": 1, "m": 1, "a": 1, "s": 1},
		Max: map[string]int{"x": 4000, "m": 4000, "a": 4000, "s": 4000},
	}
	res := wfs.LookupRange("in", 0, start)
	return res
}
