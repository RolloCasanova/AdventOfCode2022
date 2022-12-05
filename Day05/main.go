package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/RolloCasanova/AdventOfCode2022/utils/file"
)

func main() {
	input, err := file.ToStringArray("Day05/input.txt")
	if err != nil {
		panic(err)
	}

	var (
		topCrates1, topCrates2 string
		stacks, instructions   []string
		stackQty               int
		m                      map[int][]string
	)

	stacks = append(stacks, input[:8]...)
	instructions = append(instructions, input[10:]...)
	stackQty = 9

	// initialize map with the same number of stacks
	m = make(map[int][]string)
	for i := 1; i < stackQty; i++ {
		m[i] = []string{}
	}

	stacksToMap(stacks, m)

	parsedInstructions := make([][]int, len(instructions))
	for i, line := range instructions {
		move := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		matches := move.FindStringSubmatch(line)

		// convert to int and append to parsedInstructions
		parsedInstructions[i] = make([]int, 3)
		for j, match := range matches[1:] {
			parsedInstructions[i][j], err = strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
		}
	}

	for _, instruction := range parsedInstructions {
		qty := instruction[0]
		start := instruction[1]
		end := instruction[2]

		// extract cargos from start stack one by one
		for i := 0; i < qty; i++ {
			lastIndex := len(m[start]) - 1
			m[end] = append(m[end], m[start][lastIndex])
			m[start] = m[start][:lastIndex]
		}
	}

	// Part One - get the last element of all the stacks
	for i := 1; i <= stackQty; i++ {
		lastIndex := len(m[i]) - 1
		if lastIndex >= 0 {
			topCrates1 += m[i][lastIndex]
		}
	}

	// Part Two - get the last element of all the stacks using the new rearrangement procedure
	m = make(map[int][]string)
	for i := 1; i < stackQty; i++ {
		m[i] = []string{}
	}
	stacksToMap(stacks, m)

	for _, instruction := range parsedInstructions {
		qty := instruction[0]
		start := instruction[1]
		end := instruction[2]

		// extract cargos from start stack as a whole, starting from firstIndex
		firstIndex := len(m[start]) - qty
		m[end] = append(m[end], m[start][firstIndex:]...)
		m[start] = m[start][:firstIndex]
	}

	// Part Two - get the last element of all the stacks
	for i := 1; i <= stackQty; i++ {
		lastIndex := len(m[i]) - 1
		if lastIndex >= 0 {
			topCrates2 += m[i][lastIndex]
		}
	}

	fmt.Println("Part One:", topCrates1)
	fmt.Println("Part Two:", topCrates2)

}

// stacksToMap parses stacks and add them to map (from bottom to top)
// first element (bottom) of the stack is the first element in the array
func stacksToMap(stacks []string, m map[int][]string) {
	for i := len(stacks) - 1; i >= 0; i-- {
		// laxy solution assumming stacks are always one letter long
		// cargo letters are always on characters (in 0-index arrays) in positions 1, 5, 9, 13... (4n+1)
		lenStack := len(stacks[i])
		for j := 1; j < lenStack; j += 4 {
			if stacks[i][j] != ' ' {
				m[j/4+1] = append(m[j/4+1], string(stacks[i][j]))
			}
		}
	}
}
