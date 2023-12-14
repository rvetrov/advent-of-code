package utils

import "strings"

func Rotate(field []string) []string {
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
