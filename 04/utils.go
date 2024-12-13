package day04

import (
	"bufio"
	"os"
	"regexp"
)

func CharAt(str string, index int, defaultChar rune) rune {
	if index < 0 || index > len(str)-1 {
		return defaultChar
	}

	return rune(str[index])
}

func IsKeywordMatch(word string, pattern *regexp.Regexp) bool {
	return pattern.MatchString(word)
}

func ParseInput() ([]string, error) {
	file, err := os.Open("04/input.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	grid := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func SafeIndex[T any](arr []T, index int, defaultElement T) T {
	if index >= len(arr) {
		return defaultElement
	}

	return arr[index]
}
