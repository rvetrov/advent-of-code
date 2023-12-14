package utils

import (
	"fmt"
	"strings"
)

func Print(fields []string) {
	for _, line := range fields {
		fmt.Println(line)
	}
}

func buildersToLines(builders []strings.Builder) []string {
	res := make([]string, len(builders))
	for i, builder := range builders {
		res[i] = builder.String()
	}
	return res
}

func Transpose(field []string) []string {
	n, m := len(field), len(field[0])
	builders := make([]strings.Builder, m)

	for i := 0; i < n; i++ {
		for j, ch := range field[i] {
			builders[j].WriteRune(ch)
		}
	}
	return buildersToLines(builders)
}

func RotateCW(field []string) []string {
	n, m := len(field), len(field[0])
	builders := make([]strings.Builder, m)
	for j := 0; j < m; j++ {
		builder := &builders[j]
		for i := n - 1; i >= 0; i-- {
			builder.WriteByte(field[i][j])
		}
	}
	return buildersToLines(builders)
}

func RotateCCW(field []string) []string {
	n, m := len(field), len(field[0])
	builders := make([]strings.Builder, m)
	for j := 0; j < m; j++ {
		builder := &builders[m-1-j]
		for i := 0; i < n; i++ {
			builder.WriteByte(field[i][j])
		}
	}
	return buildersToLines(builders)
}
