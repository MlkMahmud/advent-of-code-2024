package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	day04Utils "github.com/MlkMahmud/advent-of-code-2024/04"
)

var (
	KEY_WORD_PATTERN = regexp.MustCompile(`(?i)^(xmas|samx)`)
)

func countMatches(grid [4]string) int {
	defaultChar := '.'
	matches := 0

	for i := 0; i < len(grid[0]); i++ {
		if day04Utils.IsKeywordMatch(grid[0][i:], KEY_WORD_PATTERN) {
			matches += 1
		}

		if word := fmt.Sprintf("%c%c%c%c", day04Utils.CharAt(grid[0], i, defaultChar), day04Utils.CharAt(grid[1], i, defaultChar), day04Utils.CharAt(grid[2], i, defaultChar), day04Utils.CharAt(grid[3], i, defaultChar)); day04Utils.IsKeywordMatch(word, KEY_WORD_PATTERN) {
			matches += 1
		}

		if word := fmt.Sprintf("%c%c%c%c", day04Utils.CharAt(grid[0], i, defaultChar), day04Utils.CharAt(grid[1], i+1, defaultChar), day04Utils.CharAt(grid[2], i+2, defaultChar), day04Utils.CharAt(grid[3], i+3, defaultChar)); day04Utils.IsKeywordMatch(word, KEY_WORD_PATTERN) {
			matches += 1
		}

		if word := fmt.Sprintf("%c%c%c%c", day04Utils.CharAt(grid[0], (len(grid[0])-1)-i, defaultChar), day04Utils.CharAt(grid[1], (len(grid[1])-2)-i, defaultChar), day04Utils.CharAt(grid[2], (len(grid[2])-3)-i, defaultChar), day04Utils.CharAt(grid[3], (len(grid[0])-4)-i, defaultChar)); day04Utils.IsKeywordMatch(word, KEY_WORD_PATTERN) {
			matches += 1
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

	for i, gridLength := 0, len(grid); i < gridLength; i++ {
		arr := [4]string{}
		arr[0] = grid[i]
		arr[1] = day04Utils.SafeIndex(grid, i+1, strings.Repeat(".", len(grid[i])))
		arr[2] = day04Utils.SafeIndex(grid, i+2, strings.Repeat(".", len(grid[i])))
		arr[3] = day04Utils.SafeIndex(grid, i+3, strings.Repeat(".", len(grid[i])))

		matches += countMatches(arr)
	}

	fmt.Printf("Ans: %d\n", matches)
}
