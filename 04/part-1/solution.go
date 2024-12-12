package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	KEY_WORD_PATTERN = regexp.MustCompile(`(?i)^(xmas|samx)`)
)

func charAt(str string, index int) string {
	if index < 0 || index > len(str)-1 {
		return "."
	}

	return string(str[index])
}

func isKeywordMatch(word string) bool {
	return KEY_WORD_PATTERN.MatchString(word)
}

func countMatches(grid []string) int {
	matches := 0

	for i := 0; i < len(grid[0]); i++ {
		if isKeywordMatch(grid[0][i:]) {
			matches += 1
		}

		if len(grid) == 4 {
			if word := fmt.Sprintf("%s%s%s%s", charAt(grid[0], i), charAt(grid[1], i), charAt(grid[2], i), charAt(grid[3], i)); isKeywordMatch(word) {
				matches += 1
			}

			if word := fmt.Sprintf("%s%s%s%s", charAt(grid[0], i), charAt(grid[1], i+1), charAt(grid[2], i+2), charAt(grid[3], i+3)); isKeywordMatch(word) {
				matches += 1
			}

			if word := fmt.Sprintf("%s%s%s%s", charAt(grid[0], (len(grid[0])-1)-i), charAt(grid[1], (len(grid[1])-2)-i), charAt(grid[2], (len(grid[2])-3)-i), charAt(grid[3], (len(grid[0])-4)-i)); isKeywordMatch(word) {
				matches += 1
			}
		}

	}

	return matches
}

func parseFile() ([]string, error) {
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

func Solution() {
	grid, err := parseFile()

	matches := 0

	if err != nil {
		log.Fatal(err)
	}

	if len(grid) < 4 {
		fmt.Printf("Ans: %d\n", matches)
		return
	}

	for i, gridLength := 0, len(grid); i < gridLength; i++ {
		if i+3 >= gridLength {
			matches += countMatches([]string{grid[i]})
		} else {
			arr := []string{grid[i], grid[i+1], grid[i+2], grid[i+3]}
			matches += countMatches(arr)
		}
	}

	fmt.Printf("Ans: %d\n", matches)
}
