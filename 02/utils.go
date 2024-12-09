package day02

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func convertStrArrToInt64Arr(arr []string) ([]int64, error) {
	result := []int64{}

	for i := 0; i < len(arr); i++ {
		value, err := strconv.ParseInt(arr[i], 10, 64)

		if err != nil {
			return nil, err
		}

		result = append(result, value)
	}

	return result, nil
}

func isSafeLevel(isIncrementing bool, previousLevel, currentLevel int64) bool {
	var diff int64

	if isIncrementing {
		diff = currentLevel - previousLevel
	} else {
		diff = previousLevel - currentLevel
	}

	return ((diff >= 1) && (diff <= 3))
}

func isSafeReport(report []int64, useDampner bool) bool {
	if len(report) < 2 {
		return true
	}

	isIncrementing := report[0] < report[1]
	reportIsSafe := true

	for level := 1; level < len(report); level++ {
		reportIsSafe = isSafeLevel(isIncrementing, report[level-1], report[level])

		if reportIsSafe {
			continue
		}

		if useDampner {
			for i := 0; i < len(report); i++ {
				subReport := []int64{}
				subReport = append(subReport, report[:i]...)
				subReport = append(subReport, report[i+1:]...)

				if isSafeReport(subReport, false) {
					reportIsSafe = true
					break
				}
			}
		}

		break
	}

	return reportIsSafe
}

func Solution(useDampener bool) (int, error) {
	file, err := os.Open("02/input.txt")

	if err != nil {
		return 0, err
	}

	defer file.Close()

	numOfSafeReports := 0

	scanner := bufio.NewScanner(file)
	scanPattern := regexp.MustCompile(`\s+`)

	for scanner.Scan() {
		entry := scanner.Text()
		entries := scanPattern.Split(entry, -1)

		report, err := convertStrArrToInt64Arr(entries)

		if err != nil {
			return 0, err
		}

		if isSafeReport(report, useDampener) {
			numOfSafeReports += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numOfSafeReports, nil
}
