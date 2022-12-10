package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RolloCasanova/AdventOfCode2022/utils/file"
)

func main() {
	input, err := file.ToStringArray("Day10/input.txt")
	if err != nil {
		panic(err)
	}

	valuesAtStart := make([]int, 241)
	valuesAtStart[0] = 0
	cycle := 1
	x := 1

	for _, line := range input {
		value := strings.Split(line, " ")

		switch len(value) {
		case 1: // noop
			valuesAtStart[cycle] = x
			cycle++
		case 2: // addx value
			n, err := strconv.Atoi(value[1])
			if err != nil {
				panic(err)
			}

			valuesAtStart[cycle] = x
			cycle++
			valuesAtStart[cycle] = x
			cycle++
			x += n
		}
	}

	// Part One - Find the signal strength during the 20th, 60th, 100th, 140th, 180th, and 220th cycles. What is the sum of these six signal strengths?
	fmt.Println("Part One:", 20*valuesAtStart[20]+60*valuesAtStart[60]+100*valuesAtStart[100]+140*valuesAtStart[140]+180*valuesAtStart[180]+220*valuesAtStart[220])

	// Part Two - Render the signal
	display := make([][]string, 6)
	for i := 0; i < 6; i++ {
		display[i] = make([]string, 40)
	}

	lit, off := "#", " "
	valuesAtStart = valuesAtStart[1:]

	for i := 0; i < len(valuesAtStart); i++ {
		pos := valuesAtStart[i] - 1
		if pos <= i%40 && i%40 <= pos+2 {
			display[i/40][i%40] = lit
		} else {
			display[i/40][i%40] = off
		}
	}

	for _, row := range display {
		fmt.Println(row)
	}
}
