package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"adventofcode.com/2023/internal/utils"
)

type Hands []*Hand

func (hs Hands) Len() int {
	return len(hs)
}

func (hs Hands) Less(i, j int) bool {
	if hs[i].Type() == hs[j].Type() {
		for k := 0; k < len(hs[i].Cards); k++ {
			if hs[i].Cards[k] != hs[j].Cards[k] {
				return hs[i].Cards[k] < hs[j].Cards[k]
			}
		}
		return false
	}
	return hs[i].Type() < hs[j].Type()
}

func (hs Hands) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

type Hand struct {
	raw string
	tp  int

	Cards []int
	Bid   int
}

func (h *Hand) Type() int {
	if h.tp == 0 {
		cntMap := map[int]int{}
		jockers := 0
		for _, strength := range h.Cards {
			if strength == 1 {
				jockers++
			} else {
				cntMap[strength]++
			}
		}

		cnt := []int{}
		for _, c := range cntMap {
			cnt = append(cnt, c)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(cnt)))

		tp := 1
		switch {
		case jockers == 5 || cnt[0]+jockers == 5:
			tp = 7
		case cnt[0]+jockers == 4:
			tp = 6
		case cnt[0]+jockers == 3 && cnt[1] == 2:
			tp = 5
		case cnt[0]+jockers == 3:
			tp = 4
		case cnt[0]+jockers == 2 && cnt[1] == 2:
			tp = 3
		case cnt[0]+jockers == 2:
			tp = 2
		}
		h.tp = tp
	}
	return h.tp
}

func charToStrength(ch rune, withJockers bool) int {
	// A, K, Q, [J], T, 9, 8, 7, 6, 5, 4, 3, 2, [J]
	switch {
	case '2' <= ch && ch <= '9':
		return int(ch - '0')
	case ch == 'T':
		return 10
	case ch == 'J':
		if withJockers {
			return 1
		} else {
			return 11
		}
	case ch == 'Q':
		return 12
	case ch == 'K':
		return 13
	case ch == 'A':
		return 14
	}
	log.Fatalf("Unknown char: %q", ch)
	return 0
}

func NewHand(str string, withJockers bool) *Hand {
	hand := Hand{}
	if n, err := fmt.Sscanf(str, "%s %d", &hand.raw, &hand.Bid); err != nil || n != 2 {
		log.Fatalf("Failed to parse hand %q. N: %q, Error: %q", str, n, err)
	}
	for _, ch := range hand.raw {
		strength := charToStrength(ch, withJockers)
		hand.Cards = append(hand.Cards, strength)
	}
	return &hand
}

func parseHands(input string, withJockers bool) Hands {
	hands := Hands{}
	for _, line := range strings.Split(input, "\n") {
		if line = strings.TrimSpace(line); len(line) > 0 {
			hands = append(hands, NewHand(line, withJockers))
		}
	}
	return hands
}

func solve(input string, withJockers bool) int {
	hands := parseHands(input, withJockers)
	sort.Sort(hands)

	res := 0
	for i, hand := range hands {
		res += (i + 1) * hand.Bid
	}
	return res
}

func solveV1(input string) int {
	return solve(input, false)
}

func solveV2(input string) int {
	return solve(input, true)
}

func main() {
	input := utils.MustReadInput("input.big")
	res := solveV2(input)
	utils.MustWriteOutput("output-v2.big", fmt.Sprint(res))
}
