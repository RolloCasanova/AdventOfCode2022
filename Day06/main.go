package main

import (
	"fmt"

	"github.com/RolloCasanova/AdventOfCode2022/utils/file"
)

func main() {
	input, err := file.ToStringArray("Day06/input.txt")
	if err != nil {
		panic(err)
	}

	m := make(map[byte]int, 4)

	sequence := input[0]

	for i := 0; i < 3; i++ {
		m[sequence[i]]++
	}

	var res1, res2 int
	for i := 3; i < len(sequence); i++ {
		m[sequence[i]]++

		if len(m) == 4 {
			res1 = i + 1
			break
		}

		if m[sequence[i-3]] >= 2 {
			m[sequence[i-3]]--
		} else {
			delete(m, sequence[i-3])
		}
	}

	m = make(map[byte]int, 14)
	for i := 0; i < 13; i++ {
		m[sequence[i]]++
	}

	for i := 14; i < len(sequence); i++ {
		m[sequence[i]]++

		if len(m) == 14 {
			res2 = i + 1
			break
		}

		if m[sequence[i-13]] >= 2 {
			m[sequence[i-13]]--
		} else {
			delete(m, sequence[i-13])
		}
	}

	// Part One: characters processed before first marker (4 characters)
	fmt.Println("Part One:", res1)

	// Part Two: characters processed before first marker (14 characters)
	fmt.Println("Part Two:", res2)

}
