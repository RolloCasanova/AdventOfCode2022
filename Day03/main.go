package main

import (
	"fmt"

	"github.com/RolloCasanova/AdventOfCode2022/utils/file"
)

func main() {
	input, err := file.ToStringArray("./Day03/input.txt")
	if err != nil {
		panic(err)
	}

	var (
		lHalf, rHalf  string
		l, sum1, sum2 int
		m1, m2        map[rune]struct{}
	)

	for _, line := range input {
		l = len(line)
		lHalf = line[:l/2]
		rHalf = line[l/2:]

		m1 = make(map[rune]struct{})
		// store all elements on first half in a map
		for _, c := range lHalf {
			m1[c] = struct{}{}
		}

		// check which element is repeated on second half, and calculate the priority
		for _, c := range rHalf {
			if _, ok := m1[c]; ok {
				sum1 += priority(c)
				break
			}
		}
	}

	// group lines by each six lines (three lines per group (3x2 is 6))
	for i := 0; i < len(input)/3; i++ {
		m1 = make(map[rune]struct{})
		// store all elements on first line
		for _, c := range input[i*3] {
			m1[c] = struct{}{}
		}

		m2 = make(map[rune]struct{})
		// store all elements repeated from m1 and second line
		for _, c := range input[i*3+1] {
			if _, ok := m1[c]; ok {
				m2[c] = struct{}{}
			}
		}

		// check which element is repeated on third line, and calculate the priority
		for _, c := range input[i*3+2] {
			if _, ok := m2[c]; ok {
				sum2 += priority(c)
				break
			}
		}
	}

	// Part One - get the sum of all priorities
	fmt.Println("Part One:", sum1)

	// Part Two - group by each three lines and get the sum of all priorities
	fmt.Println("Part Two:", sum2)
}

// priority returns the priority of a character
// a...z being 1...26, and A...Z being 27...52
func priority(r rune) int {
	var val int

	switch {
	case 'a' <= r && r <= 'z':
		val = int(r - 'a' + 1)
	case 'A' <= r && r <= 'Z':
		val = int(r-'A'+1) + 26
	}

	return val
}
