package day02

import (
	"github.com/stanislav-milchev/aoc2024/utils"
)

func isSafeDampener(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := range report {
		modedLvl := append(report[:i:i], report[i+1:]...)
		if isSafe(modedLvl) {
			return true
		}
	}

	return false
}

func Part2() int {
	// lines := utils.ReadInput("day02/input.small")
	lines := utils.ReadInput("day02/input")
	lists := MakeLists(lines)
	var sum int
	for _, list := range lists {
		if isSafeDampener(list) {
			sum++
		}
	}

	return sum
}
