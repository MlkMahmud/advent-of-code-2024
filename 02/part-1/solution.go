package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Solution() {
	file, err := os.Open("02/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	safeReports := 0

	scanner := bufio.NewScanner(file)
	scanPattern := regexp.MustCompile(`\s+`)

	for scanner.Scan() {
		entry := scanner.Text()
		entries := scanPattern.Split(entry, -1)

		report := []int64{}

		for i := 0; i < len(entries); i++ {
			value, err := strconv.ParseInt(entries[i], 10, 64)

			if err != nil {
				log.Fatal(err)
			}

			report = append(report, value)
		}

		isSafeReport := true

		for i := 0; i < len(report)-1; i++ {
			isIncrementing := report[0] < report[1]

			var diff int64

			if isIncrementing {
				diff = report[i+1] - report[i]
			} else {
				diff = report[i] - report[i+1]
			}

			if diff < 1 || diff > 3 {
				isSafeReport = false
				break
			}

		}

		if isSafeReport {
			safeReports += 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ans: %d\n", safeReports)
}
