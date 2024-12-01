package day01

import (
	"sort"
	"strconv"
	"strings"

	"github.com/stanislav-milchev/aoc2024/utils"
)

func MakeLists(lines []string) [][]int {
	var leftList []int
	var rightList []int
	for _, line := range lines {
		pair := strings.Fields(line)
		left, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}
	sort.Ints(leftList)
	sort.Ints(rightList)

	return [][]int{leftList, rightList}
}

func Part1() int {
    //lines := utils.ReadInput("day01/input.small")
    lines := utils.ReadInput("day01/input")
    lists := MakeLists(lines)
    left := lists[0]
    right := lists[1]

    var distance int
    for i := 0; i <= len(left) - 1; i++ {
        distance += utils.Abs(left[i] - right[i])
    }

    return distance
}
