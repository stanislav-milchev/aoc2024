package day01

import (
	"github.com/stanislav-milchev/aoc2024/utils"
)

func Part2() int {
	//lines := utils.ReadInput("day01/input.small")
	lines := utils.ReadInput("day01/input")
	lists := MakeLists(lines)
	left := lists[0]
	right := lists[1]

	frequency := make(map[int]int)
	for _, num := range right {
		frequency[num]++
	}

    var sum int
	for _, num := range left {
		if count, exists := frequency[num]; exists {
			sum += num * count
		}
	}

	return sum
}
