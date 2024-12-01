package utils

import (
	"bufio"
	"bytes"
	"os"
)

func ReadInput(path string) []string {
	content, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(content))
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
