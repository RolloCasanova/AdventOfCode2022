package file

import (
	"bufio"
	"os"
	"strconv"
)

// receives the input file, reads it line by line, and convert it to an array of strings
func ToStringArray(filename string) ([]string, error) {
	// open file to read
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// don't forget to close file
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var input []string

	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}

	return input, nil
}

// StringArrayToIntArray converts an array of strings to an array of integers
func StringArrayToIntArray(input []string) ([]int, error) {
	var result []int

	for _, v := range input {
		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		result = append(result, val)
	}

	return result, nil
}
