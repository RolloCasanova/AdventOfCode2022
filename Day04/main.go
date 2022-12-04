package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RolloCasanova/AdventOfCode2022/utils/file"
)

func main() {
	input, err := file.ToStringArray("Day04/input.txt")
	if err != nil {
		panic(err)
	}

	var c1, c2 int
	for _, line := range input {
		sections := strings.Split(line, ",")
		one := strings.Split(sections[0], "-")
		two := strings.Split(sections[1], "-")

		startOne, err := strconv.Atoi(one[0])
		if err != nil {
			panic(err)
		}

		endOne, err := strconv.Atoi(one[1])
		if err != nil {
			panic(err)
		}

		startTwo, err := strconv.Atoi(two[0])
		if err != nil {
			panic(err)
		}

		endTwo, err := strconv.Atoi(two[1])
		if err != nil {
			panic(err)
		}

		// fully contained
		fullyContained := (startOne <= startTwo && endTwo <= endOne) || (startTwo <= startOne && endOne <= endTwo)
		if fullyContained {
			c1++
		}

		// either fully contained or partially everlapping
		// partial overlap occurs in two cases:
		// 1. S1 <= S2 <= E1 <= E2
		// 2. S2 <= S1 <= E2 <= E1
		if fullyContained ||
			((startOne <= startTwo) && (startTwo <= endOne) && (endOne <= endTwo)) ||
			((startTwo <= startOne) && (startOne <= endTwo) && (endTwo <= endOne)) {
			c2++
		}
	}

	// Part One - count number of fully contained assignments
	fmt.Println("Part One:", c1)

	// Part Two - count number of partially contained assignments
	fmt.Println("Part Two:", c2)
}
