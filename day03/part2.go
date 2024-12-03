package day03

import (
	"regexp"
	"strconv"

	"github.com/stanislav-milchev/aoc2024/utils"
)

func parseInstructions(input string) int {
	r := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)).*?don\'t\(\)`)
	initialMatches := r.FindAllStringSubmatch(input, 1)

	re := regexp.MustCompile(`do\(\)(.*?)don\'t\(\)`)
	restMatches := re.FindAllStringSubmatch(input, -1)

	var sum int
	for _, match := range restMatches {
        result := strsToInts(match[1])
        sum += result
	}

	for _, match := range initialMatches {
		a, _ := strconv.Atoi(match[2])
		b, _ := strconv.Atoi(match[3])
		result := a * b
		sum += result
	}

	return sum
}

func Part2() int {
	//str := utils.ReadInput("day03/input.small")
	str := utils.ReadInput("day03/input")
	result := parseInstructions(str[0])
	return result
}
