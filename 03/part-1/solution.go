package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func convertStrToInt64(intStr string) (int64, error) {
	value, err := strconv.ParseInt(intStr, 10, 64)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func parseInput() ([][2]int64, error) {
	instructions := [][2]int64{}

	file, err := os.Open("03/input.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructionPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := instructionPattern.FindAllStringSubmatch(line, -1)

		for _, entry := range matches {
			multiplicand, err := convertStrToInt64(entry[1])

			if err != nil {
				return nil, err
			}

			multiplier, err := convertStrToInt64(entry[2])

			if err != nil {
				return nil, err
			}

			instructions = append(instructions, [2]int64{multiplicand, multiplier})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}

func Solution() {
	instructions, err := parseInput()
	result := 0

	if err != nil {
		log.Fatal(err)
	}

	for _, instruction := range instructions {
		result += int(instruction[0]) * int(instruction[1])
	}

	fmt.Printf("Ans: %d\n", result)
}
