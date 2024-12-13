package part2

import (
	"fmt"
	"log"
	"regexp"

	day04Utils "github.com/MlkMahmud/advent-of-code-2024/04"
)

const (
	KEY_WORD_SIZE = 3
)

var (
	KEY_WORD_PATTERN = regexp.MustCompile(`(?i)^(mas|sam)`)
)

func countMatches(grid [KEY_WORD_SIZE]string) int {
	defaultChar := '.'
	matches := 0

	for i := 0; i < len(grid[0]); i++ {
		rightDiagonalMatch := day04Utils.IsKeywordMatch(fmt.Sprintf("%c%c%c", day04Utils.CharAt(grid[0], i, defaultChar), day04Utils.CharAt(grid[1], i+1, defaultChar), day04Utils.CharAt(grid[2], i+2, defaultChar)), KEY_WORD_PATTERN)

		leftDiagonalMatch := day04Utils.IsKeywordMatch(fmt.Sprintf("%c%c%c", day04Utils.CharAt(grid[0], i+2, defaultChar), day04Utils.CharAt(grid[1], i+1, defaultChar), day04Utils.CharAt(grid[2], i, defaultChar)), KEY_WORD_PATTERN)

		if leftDiagonalMatch && rightDiagonalMatch {
			matches++
		}
	}

	return matches
}

func Solution() {
	grid, err := day04Utils.ParseInput()

	if err != nil {
		log.Fatal(err)
	}

	matches := 0

	if len(grid) < KEY_WORD_SIZE {
		fmt.Printf("Ans: %d\n", matches)
		return
	}

	for i, maxIndex := 0, len(grid)-KEY_WORD_SIZE; i <= maxIndex; i++ {
		arr := [KEY_WORD_SIZE]string{}
		arr[0] = grid[i]
		arr[1] = grid[i+1]
		arr[2] = grid[i+2]

		matches += countMatches(arr)
	}

	fmt.Printf("Ans: %d\n", matches)
}
