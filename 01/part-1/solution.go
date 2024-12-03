package part1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func parseInput() ([]int64, []int64, error) {
	var leftList []int64
	var rightList []int64

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

		if leftVal, err := strconv.ParseInt(entries[0], 10, 64); err != nil {
			return nil, nil, err
		} else {
			leftList = append(leftList, leftVal)
		}

		if rightVal, err := strconv.ParseInt(entries[1], 10, 64); err != nil {
			return nil, nil, err
		} else {
			rightList = append(rightList, rightVal)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	return leftList, rightList, nil
}

func getMaxListLength(lenListA, lenListB int) int {
	if lenListA > lenListB {
		return lenListA
	}

	return lenListB
}

func Solution() {
	leftList, rightList, err := parseInput()

	if err != nil {
		log.Fatal(err)
	}

	result, leftListLen, rightListLen := int64(0), len(leftList), len(rightList)
	maxLength := getMaxListLength(leftListLen, rightListLen)

	for i := 0; i < maxLength; i++ {
		leftListItem, rightListItem := int64(0), int64(0)

		if i < leftListLen {
			leftListItem = leftList[i]
		}

		if i < rightListLen {
			rightListItem = rightList[i]
		}

		result += int64(math.Abs(float64(leftListItem) - float64(rightListItem)))
	}

	fmt.Printf("Ans: %d\n", result)
}
