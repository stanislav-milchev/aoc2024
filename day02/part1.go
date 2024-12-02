package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stanislav-milchev/aoc2024/utils"
)

func strsToInts(strs []string) ([]int, error) {
	ints := make([]int, len(strs))
	for i, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("you are a failure: %w", err)
		}
		ints[i] = num
	}
	return ints, nil
}

func MakeLists(lines []string) [][]int {
	var intLvls [][]int
	for _, line := range lines {
		levels := strings.Fields(line)
		lvls, err := strsToInts(levels)
		if err != nil {
			panic(err)
		}
		intLvls = append(intLvls, lvls)

	}
	return intLvls
}

func isSafe(report []int) bool {
	inc := report[0] < report[1]

	var check func([]int, bool) bool

	check = func(report []int, inc bool) bool {
		if len(report) < 2 {
			return true
		}
		if inc && report[0] > report[1] {
			return false
		}

		if !inc && report[0] < report[1] {
			return false
		}

		if utils.Abs(report[0]-report[1]) < 1 || utils.Abs(report[0]-report[1]) > 3 {
			return false
		}
		return check(report[1:], inc)
	}

	return check(report, inc)
}

func Part1() int {
	// lines := utils.ReadInput("day02/input.small")
	lines := utils.ReadInput("day02/input")
	lists := MakeLists(lines)
	var sum int
	for _, list := range lists {
		if isSafe(list) {
			sum++
		}
	}

	return sum
}
