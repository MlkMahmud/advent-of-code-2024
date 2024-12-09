package part2

import (
	"fmt"
	"log"

	day02 "github.com/MlkMahmud/advent-of-code-2024/02"
)

func Solution() {
	safeReports, err := day02.Solution(true)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Ans: %d\n", safeReports)
}
