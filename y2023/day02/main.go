package day02

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type CubeSet struct {
	Reds   int
	Greens int
	Blues  int
}

func NewCubeSet(line string) *CubeSet {
	cubeSet := &CubeSet{}

	//1 red, 2 green, 6 blue
	for _, part := range strings.Split(line, ",") {
		numStr, color, _ := strings.Cut(strings.TrimSpace(part), " ")
		switch color {
		case "red":
			cubeSet.Reds, _ = strconv.Atoi(numStr)
		case "green":
			cubeSet.Greens, _ = strconv.Atoi(numStr)
		case "blue":
			cubeSet.Blues, _ = strconv.Atoi(numStr)
		}
	}
	return cubeSet
}

func (cs *CubeSet) LTE(other *CubeSet) bool {
	return cs.Reds <= other.Reds && cs.Greens <= other.Greens && cs.Blues <= other.Blues
}

func (cs *CubeSet) Power() int {
	return cs.Reds * cs.Greens * cs.Blues
}

type Game struct {
	ID       int
	reveales []CubeSet
}

func NewGame(line string) *Game {
	game := &Game{}
	gameStr, revealsStr, _ := strings.Cut(line, ":")
	if n, err := fmt.Sscanf(gameStr, "Game %d", &game.ID); n != 1 || err != nil {
		log.Fatalf("Faield to parse. Error: %q. Line: %q", err, line)
	}
	for _, revealStr := range strings.Split(revealsStr, ";") {
		game.reveales = append(game.reveales, *NewCubeSet(revealStr))
	}
	return game
}

func (g *Game) MinCubeSet() *CubeSet {
	max := &CubeSet{0, 0, 0}
	for _, reveal := range g.reveales {
		if max.Reds < reveal.Reds {
			max.Reds = reveal.Reds
		}
		if max.Greens < reveal.Greens {
			max.Greens = reveal.Greens
		}
		if max.Blues < reveal.Blues {
			max.Blues = reveal.Blues
		}
	}
	return max
}

func SolveV1(input string) int {
	res := 0
	lines := strings.Split(input, "\n")

	target := &CubeSet{Reds: 12, Greens: 13, Blues: 14}
	for _, line := range lines {
		game := NewGame(line)
		if game.MinCubeSet().LTE(target) {
			res += game.ID
		}
	}
	return res
}

func SolveV2(input string) int {
	res := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		game := NewGame(line)
		res += game.MinCubeSet().Power()
	}
	return res
}
