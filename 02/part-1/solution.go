package part1

import (
	"fmt"
	"log"

	day02 "github.com/MlkMahmud/advent-of-code-2024/02"
)

func Solution() {
	safeReports, err := day02.Solution(false)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Ans: %d\n", safeReports)
}
