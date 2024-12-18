package main

import (
	"fmt"
	"time"

	"github.com/stanislav-milchev/aoc2024/day01"
	"github.com/stanislav-milchev/aoc2024/day02"
	"github.com/stanislav-milchev/aoc2024/day03"
)

func measureAndPrint[T any](label string, fn func() T) {
	// return
	startTime := time.Now()
	res := fn()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("%s: %-15v\t(%v)\n", label, res, elapsed)
}

func main() {
	measureAndPrint("D01P1", day01.Part1)
	measureAndPrint("D01P2", day01.Part2)
	measureAndPrint("D02P1", day02.Part1)
	measureAndPrint("D02P2", day02.Part2)
	measureAndPrint("D03P1", day03.Part1)
	measureAndPrint("D03P2", day03.Part2)
}
