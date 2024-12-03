package day03

import (
	"regexp"
	"strconv"

	"github.com/stanislav-milchev/aoc2024/utils"
)

func strsToInts(input string) int {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
    matches := r.FindAllStringSubmatch(input, -1)

    var sum int
    for _, match := range matches {
        a, _ := strconv.Atoi(match[1])
        b, _ := strconv.Atoi(match[2])
        result := a * b
        sum += result
    }

    return sum
}

func Part1() int {
	//str := utils.ReadInput("day02/input.small")
	str := utils.ReadInput("day03/input")
    result := strsToInts(str[0])
    return result

}
