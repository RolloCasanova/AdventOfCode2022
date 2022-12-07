package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RolloCasanova/AdventOfCode2022/utils/file"
)

type dir struct {
	name     string
	size     int
	children map[string]*dir
	parent   *dir
}

func main() {
	input, err := file.ToStringArray("Day07/input.txt")
	if err != nil {
		panic(err)
	}

	currentDirectory := &dir{
		name: "/",
	}

	// separate commands and output
	for _, line := range input[1:] {
		words := strings.Split(line, " ")
		switch words[0] {
		case "$": // is a command
			switch words[1] {
			case "cd":
				switch words[2] {
				case "..":
					currentDirectory = currentDirectory.parent
				default:
					currentDirectory = currentDirectory.children[words[2]]
				}
				// case "ls": nothing to do here
			}

		case "dir": // is a  directory
			if currentDirectory.children == nil {
				currentDirectory.children = make(map[string]*dir)
			}

			currentDirectory.children[words[1]] = &dir{
				name:   words[1],
				parent: currentDirectory,
			}

		default: // is a file
			filesize, err := strconv.Atoi(words[0])
			if err != nil {
				panic(err)
			}

			if currentDirectory.children == nil {
				currentDirectory.children = make(map[string]*dir)
			}

			currentDirectory.children[words[1]] = &dir{
				name:   words[1],
				size:   filesize,
				parent: currentDirectory,
			}
		}
	}

	// go to root
	for currentDirectory.parent != nil {
		currentDirectory = currentDirectory.parent
	}

	// iterate over all directories and calculate the size of each folder
	setDirSize(currentDirectory)

	// go to root
	for currentDirectory.parent != nil {
		currentDirectory = currentDirectory.parent
	}

	// Part One: total size of directories up to 100000
	var sum int
	sumSize(currentDirectory, &sum)

	fmt.Println("Part One:", sum)

	// go to root
	for currentDirectory.parent != nil {
		currentDirectory = currentDirectory.parent
	}

	// Part Two: Delete a file to free up at least 30000000 space from a 70000000 filesystem
	toDelete := 30000000 - (70000000 - currentDirectory.size)
	minDirSize := 1<<63 - 1

	// find the smallest dir with at least freeSpace size
	findSmallestDir(currentDirectory, toDelete, &minDirSize)

	fmt.Println("Part Two:", minDirSize)

}

func findSmallestDir(dir *dir, toDelete int, minDirSize *int) {
	if dir.name != "/" {
		if dir.size >= toDelete && dir.children != nil && dir.size < *minDirSize {
			*minDirSize = dir.size
		}
	}
	for _, child := range dir.children {
		findSmallestDir(child, toDelete, minDirSize)
	}
}

func setDirSize(dir *dir) int {
	if dir.children == nil {
		return dir.size
	}

	for _, child := range dir.children {
		dir.size += setDirSize(child)
	}

	return dir.size
}

func sumSize(dir *dir, sum *int) {
	if dir.size <= 100000 && dir.children != nil {
		*sum += dir.size
	}

	for _, child := range dir.children {
		sumSize(child, sum)
	}
}
