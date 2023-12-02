package utils

import "os"

func MustReadInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(input)
}

func MustWriteOutput(fileName string, text string) {
	os.WriteFile(fileName, []byte(text), 0644)
}
