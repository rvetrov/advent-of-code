package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"adventofcode.com/internal/utils"
)

func rotate(field []string) []string {
	builders := make([]strings.Builder, len(field[0]))

	for i := 0; i < len(field); i++ {
		for j, ch := range field[i] {
			builders[j].WriteRune(ch)
		}
	}

	res := []string{}
	for _, builder := range builders {
		res = append(res, builder.String())
	}
	return res
}

func findHorizontalReflectionLines(field []string) []int {
	res := []int{}
	for i := 1; i < len(field); i++ {
		found := true
		for j := 0; i+j < len(field) && i-1-j >= 0; j++ {
			if field[i+j] != field[i-1-j] {
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

func findReflectionLines(field []string) ([]int, []int) {
	rowRefl := findHorizontalReflectionLines(field)
	rotated := rotate(field)
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

func solveV1(input string) int {
	fields := utils.SplitByEmptyLine(input)
	res := 0
	for _, field := range fields {
		if rr, cr := findReflectionLines(field); len(rr) == 1 || len(cr) == 1 {
			res += encodeReflections(rr, cr, nil, nil)
		} else {
			log.Fatalln("No reflections in case:", field)
		}
	}
	return res
}

func alternativeReflection(field []string, origRows, origCols []int) ([]int, []int) {
	for i := 0; i < len(field); i++ {
		origLine := field[i]
		lineBytes := []byte(origLine)
		for j, ch := range lineBytes {
			if ch == '.' {
				lineBytes[j] = '#'
			} else {
				lineBytes[j] = '.'
			}
			field[i] = string(lineBytes)

			if rr, cr := findReflectionLines(field); encodeReflections(rr, cr, origRows, origCols) > 0 {
				return rr, cr
			}

			lineBytes[j] = ch
		}

		field[i] = origLine
	}
	fmt.Println(field)
	log.Fatalln("Failed to found")
	return nil, nil
}

func solveV2(input string) int {
	fields := utils.SplitByEmptyLine(input)
	res := 0
	for _, field := range fields {
		if origRows, origCols := findReflectionLines(field); len(origRows) == 1 || len(origCols) == 1 {
			rr, cr := alternativeReflection(field, origRows, origCols)
			res += encodeReflections(rr, cr, origRows, origCols)
		} else {
			log.Fatalln("No reflections in case:", field)
		}
	}
	return res
}

func main() {
	input := utils.MustReadInput("input.big")
	res := solveV2(input)
	utils.MustWriteOutput("output-v2.big", fmt.Sprint(res))
}
