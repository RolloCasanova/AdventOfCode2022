package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	// open file to read
	file, err := os.Open("./Day01/input.txt")
	if err != nil {
		panic(err)
	}

	// don't forget to close file
	defer file.Close()

	var calories []int
	var num, sum int

	// read the file line by line, exit when EOF
	for {
		_, err := fmt.Fscanf(file, "%d", &num)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			calories = append(calories, sum)
			sum = 0
			continue
		}

		sum += num
	}

	// sort calories' slice
	sort.IntSlice(calories).Sort()

	l := len(calories)

	// Part One - get max value of calories
	fmt.Println("Part One:", calories[l-1])

	// Part Two - get sum of max three values of calories
	fmt.Println("Part Two:", calories[l-1]+calories[l-2]+calories[l-3])
}
