package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/RolloCasanova/AdventOfCode2022/utils/file"
)

func main() {
	input, err := file.ToStringArray("./Day01/input.txt")
	if err != nil {
		panic(err)
	}

	var (
		calories []int
		sum      int
	)

	// read line by line, converting and adding values
	for _, line := range input {
		if line == "" {
			calories = append(calories, sum)
			sum = 0
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		sum += val
	}

	// sort calories' slice
	sort.IntSlice(calories).Sort()

	l := len(calories)

	// Part One - get max value of calories
	fmt.Println("Part One:", calories[l-1])

	// Part Two - get sum of max three values of calories
	fmt.Println("Part Two:", calories[l-1]+calories[l-2]+calories[l-3])
}
