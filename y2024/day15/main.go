package day15

import (
	"strings"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

const (
	freeChar     = '.'
	robotChar    = '@'
	boxChar      = 'O'
	boxLeftChar  = '['
	boxRightChar = ']'
	wallChar     = '#'
)

func moveToDir(move byte) grid.Direction {
	switch move {
	case '<':
		return grid.Left
	case '>':
		return grid.Right
	case '^':
		return grid.Up
	case 'v':
		return grid.Down
	default:
		panic(move)
	}
}

func makeSimpleMove(gr grid.Grid, pos grid.Position, dir grid.Direction) grid.Position {
	freePos := pos.Add(dir)
	freePosChar, contains := gr.At(freePos)
	for contains && (freePosChar == boxChar || freePosChar == boxLeftChar || freePosChar == boxRightChar) {
		freePos = freePos.Add(dir)
		freePosChar, contains = gr.At(freePos)
	}
	if !contains || freePosChar == wallChar {
		return pos
	}

	reversedDir := dir.Reversed()
	lastPos := freePos
	for lastPos != pos {
		prevPos := lastPos.Add(reversedDir)
		prevPosCh, _ := gr.At(prevPos)
		gr.SetAt(lastPos, prevPosCh)
		lastPos = prevPos
	}
	gr.SetAt(pos, freeChar)
	return pos.Add(dir)
}

func makeCascadeVerticalMove(gr grid.Grid, pos grid.Position, dir grid.Direction) grid.Position {
	var levels [][]grid.Position
	var levelToMove []grid.Position
	levelToMove = append(levelToMove, pos)

	for len(levelToMove) > 0 {
		levels = append(levels, levelToMove)
		var newLevel []grid.Position

		for _, posToMove := range levelToMove {
			candPos := posToMove.Add(dir)
			candCh, _ := gr.At(candPos)

			switch candCh {
			case wallChar:
				return pos
			case boxLeftChar:
				newLevel = append(newLevel, candPos, candPos.Add(grid.Right))
			case boxRightChar:
				if len(newLevel) == 0 || newLevel[len(newLevel)-1] != candPos {
					newLevel = append(newLevel, candPos.Add(grid.Left), candPos)
				}
			}
		}

		levelToMove = newLevel
	}

	for i := len(levels) - 1; i >= 0; i-- {
		for _, levelPos := range levels[i] {
			ch, _ := gr.At(levelPos)
			gr.SetAt(levelPos.Add(dir), ch)
			gr.SetAt(levelPos, freeChar)
		}
	}

	return pos.Add(dir)
}

func sumOfBoxesCoordinates(gr grid.Grid) int {
	res := 0
	for pos := gr.Start(); gr.Contains(pos); pos = gr.Next(pos) {
		ch, _ := gr.At(pos)
		if ch == boxChar || ch == boxLeftChar {
			res += pos.Row*100 + pos.Col
		}
	}
	return res
}

func widenMap(lines []string) []string {
	var res []string
	for _, line := range lines {
		var resLine []byte
		for _, ch := range line {
			switch ch {
			case freeChar:
				resLine = append(resLine, freeChar, freeChar)
			case wallChar:
				resLine = append(resLine, wallChar, wallChar)
			case robotChar:
				resLine = append(resLine, robotChar, freeChar)
			case boxChar:
				resLine = append(resLine, boxLeftChar, boxRightChar)
			}
		}
		res = append(res, string(resLine))
	}
	return res
}

func SolveV1(input string) int {
	blocks := utils.SplitByEmptyLine(input)
	gr := grid.New(blocks[0])
	robotPos, _ := gr.FindPosition(robotChar)

	for _, move := range []byte(strings.Join(blocks[1], "")) {
		dir := moveToDir(move)
		robotPos = makeSimpleMove(gr, robotPos, dir)
	}

	//grid.Print(gr)
	return sumOfBoxesCoordinates(gr)
}

func SolveV2(input string) int {
	blocks := utils.SplitByEmptyLine(input)
	gr := grid.New(widenMap(blocks[0]))
	robotPos, _ := gr.FindPosition(robotChar)

	for _, move := range []byte(strings.Join(blocks[1], "")) {
		dir := moveToDir(move)
		if dir == grid.Left || dir == grid.Right {
			robotPos = makeSimpleMove(gr, robotPos, dir)
		} else {
			robotPos = makeCascadeVerticalMove(gr, robotPos, dir)
		}
	}

	//grid.Print(gr)
	return sumOfBoxesCoordinates(gr)
}
