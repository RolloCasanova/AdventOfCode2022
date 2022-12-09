package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RolloCasanova/AdventOfCode2022/utils/file"
)

type (
	// pos represents a x,y pair coordinate, where x is the row and y is the column
	// x to the right and y to the bottom are positives, where x to the left and y to the top are negatives
	pos struct {
		x int
		y int
	}

	motion struct {
		direction string
		qty       int
	}
)

var (
	exists struct{}
	m      map[pos]struct{} = make(map[pos]struct{})

	// tailLocation represents the relative direction tail is in comparison with head
	// head is always the center of the matrix (5), and at the beginning tail is also at the center
	// 1 2 3
	// 4 5 6
	// 7 8 9
	tailLocation int

	// tailLocations []int
)

func main() {
	input, err := file.ToStringArray("Day09/input.txt")
	if err != nil {
		panic(err)
	}

	var (
		num   int
		parts []string
	)

	motions := []motion{}
	for _, line := range input {
		parts = strings.Split(line, " ")
		num, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		motions = append(motions, motion{parts[0], num})
	}

	var head, tail pos
	m[tail] = exists
	tailLocation = 5

	for _, mot := range motions {
		for q := 0; q < mot.qty; q++ {
			defineTailMovement(mot.direction, &tailLocation, &head, &tail, true)
		}
	}

	// Part One: Number of distinct locations visited by the tail
	fmt.Println("Part One:", len(m))

	// // Part Two: Number of distinct locations visited by the tail in a 10 knots rope
	// m := make(map[pos]struct{})
	// heads := make([]pos, 10)
	// tailLocations = make([]int, 9)
	// for i := 0; i < len(tailLocations); i++ {
	// 	tailLocations[i] = 5
	// }
	// tail.x, tail.y = 0, 0

	// m[tail] = exists

	// for _, mot := range motions {
	// 	for q := 0; q < mot.qty; q++ {
	// 		for i := 0; i < len(heads)-1; i++ {
	// 			if i > q {
	// 				break
	// 			}
	// 			defineTailMovement(mot.direction, &tailLocations[i], &heads[i], &heads[i+1], i+1 == len(heads))
	// 		}
	// 	}
	// }

	// fmt.Println("Part Two:", len(m))
}

func defineTailMovement(direction string, tailPos *int, h, t *pos, trackTail bool) {
	switch direction {
	case "U":
		h.y--
		switch *tailPos {
		case 1, 2, 3, 4, 6:
			*tailPos += 3
		default: // tail will move to the space head was in the previous step
			*tailPos = 8
		}

	case "D":
		h.y++
		switch *tailPos {
		case 4, 6, 7, 8, 9:
			*tailPos -= 3
		default: // tail will move to the space head was in the previous step
			*tailPos = 2
		}

	case "L":
		h.x--
		switch *tailPos {
		case 1, 2, 4, 7, 8:
			*tailPos++
		default: // tail will move to the space head was in the previous step
			*tailPos = 6
		}

	case "R":
		h.x++
		switch *tailPos {
		case 2, 3, 6, 8, 9:
			*tailPos--
		default: // tail will move to the space head was in the previous step
			*tailPos = 4
		}
	}

	if trackTail {
		switch *tailPos {
		case 1:
			t.x, t.y = h.x-1, h.y-1
		case 2:
			t.x, t.y = h.x, h.y-1
		case 3:
			t.x, t.y = h.x+1, h.y-1
		case 4:
			t.x, t.y = h.x-1, h.y
		case 5:
			t.x, t.y = h.x, h.y
		case 6:
			t.x, t.y = h.x+1, h.y
		case 7:
			t.x, t.y = h.x-1, h.y+1
		case 8:
			t.x, t.y = h.x, h.y+1
		case 9:
			t.x, t.y = h.x+1, h.y+1
		}

		m[*t] = exists
	}

}
