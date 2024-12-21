package day21

import (
	"strconv"
	"strings"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

type Move struct {
	from byte
	to   byte
}

type MoveCosts map[Move]int

var (
	arrowButtons = []byte{'^', 'A', '<', 'v', '>'}
	dirToArrow   = map[grid.Direction]byte{
		grid.Up:    '^',
		grid.Down:  'v',
		grid.Right: '>',
		grid.Left:  '<',
	}

	arrowsGrid = grid.New([]string{
		" ^A",
		"<v>",
	})
	numsGrid = grid.New([]string{
		"789",
		"456",
		"123",
		" 0A",
	})
)

type bfsItem struct {
	usedArrow byte
	pos       grid.Position
}

func bfs(gr grid.Grid, start grid.Position, arrowCosts, resultCosts MoveCosts) {
	startItem := bfsItem{usedArrow: 'A', pos: start}
	startCh, _ := gr.At(start)
	costs := make(map[bfsItem]int)
	costs[startItem] = 0
	q := []bfsItem{startItem}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		curPosCost := costs[cur]
		curPosCh, _ := gr.At(cur.pos)

		costToCurPos := curPosCost + arrowCosts[Move{from: cur.usedArrow, to: 'A'}]
		moveToCurPos := Move{from: startCh, to: curPosCh}
		if knownCost, exists := resultCosts[moveToCurPos]; !exists || knownCost > costToCurPos {
			resultCosts[moveToCurPos] = costToCurPos
		}

		for _, dir := range grid.FourSides {
			nextPos := cur.pos.Add(dir)
			if ch, inside := gr.At(nextPos); !inside || ch == ' ' {
				continue
			}

			arrow := dirToArrow[dir]
			cost := curPosCost + arrowCosts[Move{from: cur.usedArrow, to: arrow}]
			nextItem := bfsItem{usedArrow: arrow, pos: nextPos}
			knownCost, exists := costs[nextItem]
			if !exists {
				q = append(q, nextItem)
				costs[nextItem] = cost
			} else if knownCost > cost {
				costs[nextItem] = cost
			}
		}
	}
}

func calcAllMoveCosts(gr grid.Grid, arrowCosts MoveCosts) MoveCosts {
	newCosts := make(MoveCosts)
	for startPos := gr.First(); gr.Contains(startPos); startPos = gr.Next(startPos) {
		if startCh, _ := gr.At(startPos); startCh != ' ' {
			bfs(gr, startPos, arrowCosts, newCosts)
		}
	}
	return newCosts
}

func calcArrowCosts(chainLength int) MoveCosts {
	costs := make(MoveCosts)
	for _, from := range arrowButtons {
		for _, to := range arrowButtons {
			costs[Move{from, to}] = 1
		}
	}

	for range chainLength {
		costs = calcAllMoveCosts(arrowsGrid, costs)
	}
	return costs
}

func solve(input string, chainLength int) int {
	arrowCosts := calcArrowCosts(chainLength)
	numsCosts := calcAllMoveCosts(numsGrid, arrowCosts)

	complexity := 0
	for _, line := range utils.NonEmptyLines(input) {
		var prevCh byte = 'A'

		numStr, _ := strings.CutSuffix(line, "A")
		num, _ := strconv.Atoi(numStr)
		cost := 0
		for _, ch := range []byte(line) {
			cost += numsCosts[Move{from: prevCh, to: ch}]
			prevCh = ch
		}

		complexity += num * cost
	}
	return complexity
}

func SolveV1(input string) int {
	return solve(input, 2)
}

func SolveV2(input string) int {
	return solve(input, 25)
}
