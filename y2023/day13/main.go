package day13

import (
	"log"
	"slices"

	"adventofcode.com/internal/grid"
	"adventofcode.com/internal/utils"
)

func findHorizontalReflectionLines(gr grid.Grid) []int {
	var res []int
	lines := gr.Lines()
	for i := 1; i < len(lines); i++ {
		found := true
		for j := 0; i+j < len(lines) && i-1-j >= 0; j++ {
			if lines[i+j] != lines[i-1-j] {
				found = false
				break
			}
		}
		if found {
			res = append(res, i)
		}
	}
	return res
}

func findReflectionLines(gr grid.Grid) ([]int, []int) {
	rowRefl := findHorizontalReflectionLines(gr)
	rotated := grid.Transpose(gr)
	colRefl := findHorizontalReflectionLines(rotated)
	return rowRefl, colRefl
}

func encodeReflections(rows, cols, rowExcludes, colExcludes []int) int {
	res := 0
	for _, row := range rows {
		if !slices.Contains(rowExcludes, row) {
			res += row * 100
		}
	}
	for _, col := range cols {
		if !slices.Contains(colExcludes, col) {
			res += col
		}
	}
	return res
}

func SolveV1(input string) int {
	grids := utils.SplitByEmptyLine(input)
	res := 0
	for _, lines := range grids {
		gr := grid.New(lines)
		if rr, cr := findReflectionLines(gr); len(rr) == 1 || len(cr) == 1 {
			res += encodeReflections(rr, cr, nil, nil)
		} else {
			log.Fatalln("No reflections in case:", gr)
		}
	}
	return res
}

func alternativeReflection(gr grid.Grid, origRows, origCols []int) ([]int, []int) {
	lines := gr.Lines()
	for i := 0; i < len(lines); i++ {
		origLine := lines[i]
		lineBytes := []byte(origLine)
		for j, ch := range lineBytes {
			if ch == '.' {
				lineBytes[j] = '#'
			} else {
				lineBytes[j] = '.'
			}
			lines[i] = string(lineBytes)

			if rr, cr := findReflectionLines(grid.New(lines)); encodeReflections(rr, cr, origRows, origCols) > 0 {
				return rr, cr
			}

			lineBytes[j] = ch
		}

		lines[i] = origLine
	}
	grid.Print(gr)
	log.Fatalln("Failed to found")
	return nil, nil
}

func SolveV2(input string) int {
	grids := utils.SplitByEmptyLine(input)
	res := 0
	for _, lines := range grids {
		gr := grid.New(lines)
		if origRows, origCols := findReflectionLines(gr); len(origRows) == 1 || len(origCols) == 1 {
			rr, cr := alternativeReflection(gr, origRows, origCols)
			res += encodeReflections(rr, cr, origRows, origCols)
		} else {
			log.Fatalln("No reflections in case:", gr)
		}
	}
	return res
}
