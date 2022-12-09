package main

import (
	"fmt"

	"github.com/RolloCasanova/AdventOfCode2022/utils/file"
)

func main() {
	input, err := file.ToStringArray("Day08/input.txt")
	if err != nil {
		panic(err)
	}

	// parse input to an int matrix
	forest, err := file.StringArrayToIntMatrix(input)
	if err != nil {
		panic(err)
	}

	// create map to store the XX-YY position of each tree that is visible from any position
	// the key is the XX-YY position and the value is just an empty struct
	visibleTrees := make(map[string]struct{})

	// check which trees are visible from each position, being 0 the lowest height and 9 the highest height

	var minHeight int

	// by rows
	for i, row := range forest {
		// from left
		minHeight = -1
		for j, tree := range row {
			if tree > minHeight {
				visibleTrees[fmt.Sprintf("%02d-%02d", i, j)] = struct{}{}
				minHeight = tree
				if minHeight == 9 { // no more trees can be seen further a 9
					break
				}
			}
		}

		// from right
		minHeight = -1
		for j := len(row) - 1; j >= 0; j-- {
			tree := row[j]
			if tree > minHeight {
				visibleTrees[fmt.Sprintf("%02d-%02d", i, j)] = struct{}{}
				minHeight = tree
				if minHeight == 9 { // no more trees can be seen further a 9
					break
				}
			}
		}
	}

	// by columns
	for j := 0; j < len(forest); j++ {
		// from top
		minHeight = -1
		for i := 0; i < len(forest); i++ {
			tree := forest[i][j]
			if tree > minHeight {
				visibleTrees[fmt.Sprintf("%02d-%02d", i, j)] = struct{}{}
				minHeight = tree
				if minHeight == 9 { // no more trees can be seen further a 9
					break
				}
			}
		}

		// from bottom
		minHeight = -1
		for i := len(forest) - 1; i >= 0; i-- {
			tree := forest[i][j]
			if tree > minHeight {
				visibleTrees[fmt.Sprintf("%02d-%02d", i, j)] = struct{}{}
				minHeight = tree
				if minHeight == 9 { // no more trees can be seen further a 9
					break
				}
			}
		}
	}

	// Part One: number of trees that are visible from any position
	fmt.Println("Part one:", len(visibleTrees))

	// Part Two: maximum scenic score
	maxScore := calculateScore(forest)

	fmt.Println("Part two:", maxScore)

}

func calculateScore(forest [][]int) int {
	var tree, score, scoreLeft, scoreRight, scoreUp, scoreDown, maxScore int

	for i := range forest {
		for j := range forest[i] {
			scoreLeft, scoreRight, scoreUp, scoreDown = 0, 0, 0, 0
			tree = forest[i][j]

			// calculate score to the left
			for x := j - 1; x >= 0; x-- {
				if forest[i][x] < tree {
					scoreLeft++
					continue
				}

				if forest[i][x] == tree {
					scoreLeft++
				}
				break

			}

			if scoreLeft == 0 {
				continue
			}

			// calculate score to the right
			for x := j + 1; x < len(forest[i]); x++ {
				if forest[i][x] < tree {
					scoreRight++
					continue
				}

				if forest[i][x] == tree {
					scoreRight++
				}

				break

			}

			if scoreRight == 0 {
				continue
			}

			// calculate score to the top
			for y := i - 1; y >= 0; y-- {
				if forest[y][j] < tree {
					scoreUp++
					continue
				}

				if forest[y][j] == tree {
					scoreUp++
				}
				break
			}

			if scoreUp == 0 {
				continue
			}

			// calculate score to the bottom
			for y := i + 1; y < len(forest); y++ {
				if forest[y][j] < tree {
					scoreDown++
					continue
				}

				if forest[y][j] == tree {
					scoreDown++
				}
				break

			}

			if scoreDown == 0 {
				continue
			}

			score = scoreLeft * scoreRight * scoreUp * scoreDown

			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}
