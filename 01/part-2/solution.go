package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parseInput() ([]int64, map[int64]int64, error) {
	var leftList []int64
	var rightListFrequencyMap map[int64]int64 = map[int64]int64{}

	file, err := os.Open("01/input.txt")

	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	entryPattern := regexp.MustCompile(`\s+`)

	for scanner.Scan() {
		entry := scanner.Text()
		entries := entryPattern.Split(entry, -1)

		leftVal, err := strconv.ParseInt(entries[0], 10, 64)

		if err != nil {
			return nil, nil, err
		}

		leftList = append(leftList, leftVal)

		rightVal, err := strconv.ParseInt(entries[1], 10, 64)

		if err != nil {
			return nil, nil, err
		}

		if val, ok := rightListFrequencyMap[rightVal]; ok {
			rightListFrequencyMap[rightVal] = val + 1
		} else {
			rightListFrequencyMap[rightVal] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftList, rightListFrequencyMap, nil
}

func Solution() {
	leftList, rightListFrequencyMap, err := parseInput()

	if err != nil {
		log.Fatal(err)
	}

	result := int64(0)

	for _, id := range leftList {
		if val, ok := rightListFrequencyMap[id]; ok {
			result += id * val
		} else {
			result += id * 0
		}
	}

	fmt.Printf("Ans: %d\n", result)
}
