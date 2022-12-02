package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	rock     = "X"
	paper    = "Y"
	scissors = "Z"

	lose = 0
	draw = 3
	win  = 6
)

var (
	// Selecting Rock grants you 1 point, Paper grants you 2 points, and Scissors grants you 3 points.
	RoundMap = map[string]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}

	// Opponent:	A = Rock	B = Paper	C = Scissors
	// You: 		X = Rock	Y = Paper	Z = Scissors
	// Score:		6 = Win		3 = Draw	0 = Lose
	Outcome1Map = map[string]int{
		"A X": draw + RoundMap[rock],
		"A Y": win + RoundMap[paper],
		"A Z": lose + RoundMap[scissors],
		"B X": lose + RoundMap[rock],
		"B Y": draw + RoundMap[paper],
		"B Z": win + RoundMap[scissors],
		"C X": win + RoundMap[rock],
		"C Y": lose + RoundMap[paper],
		"C Z": draw + RoundMap[scissors],
	}

	// Opponent:	A = Rock	B = Paper	C = Scissors
	// You: 		X = Lose	Y = Draw	Z = Win
	// Score:		6 = Win		3 = Draw	0 = Lose
	Outcome2Map = map[string]int{
		"A X": lose + RoundMap[scissors],
		"A Y": draw + RoundMap[rock],
		"A Z": win + RoundMap[paper],
		"B X": lose + RoundMap[rock],
		"B Y": draw + RoundMap[paper],
		"B Z": win + RoundMap[scissors],
		"C X": lose + RoundMap[paper],
		"C Y": draw + RoundMap[scissors],
		"C Z": win + RoundMap[rock],
	}
)

func main() {
	// open file to read
	file, err := os.Open("./Day02/input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var input []string

	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}

	// don't forget to close file
	defer file.Close()

	// store strategy score
	var score1, score2 int

	for _, round := range input {
		score1 += Outcome1Map[round]
		score2 += Outcome2Map[round]
	}

	// Part One - get total score
	fmt.Println("Part One:", score1)

	// Part Two - get total score with new strategy
	fmt.Println("Part Two:", score2)

}
