package main

import (
	"flag"
	"fmt"
	"log"

	dayOnePtOne "github.com/MlkMahmud/advent-of-code-2024/01/part-1"
	dayOnePtTwo "github.com/MlkMahmud/advent-of-code-2024/01/part-2"
	dayTwoPtOne "github.com/MlkMahmud/advent-of-code-2024/02/part-1"
)

var commands = map[string]func(){
	"day-1-part-1": dayOnePtOne.Solution,
	"day-1-part-2": dayOnePtTwo.Solution,
	"day-2-part-1": dayTwoPtOne.Solution,
}

func validateOptions(day, part int) bool {
	if day < 1 || day > 25 {
		return false
	}

	if part != 1 && part != 2 {
		return false
	}

	return true
}

func printHelpMessage() {
	log.Fatalln(("Usage:\n go run main.go --day <1-25> --part <1|2>\n\nOptions:\n --day\tSpecifies the day of the month. Must be a number between 1 and 25\n --part\tSpecifies the part. Must be either 1 or 2.\n\nNote:\n Both arguments are mandatory. Ensure you provide valid values for both \"--day\" and \"--part\"."))
}

func main() {
	log.SetFlags(0)

	var day int
	var part int

	flag.IntVar(&day, "day", 0, "Specifies the day of the month. Must be a number between 1 and 25")
	flag.IntVar(&part, "part", 0, "Specifies the part. Must be either 1 or 2.")

	flag.Parse()

	if isValidOptions := validateOptions(day, part); !isValidOptions {
		printHelpMessage()
	}

	commandKey := fmt.Sprintf("day-%d-part-%d", day, part)

	if command, ok := commands[commandKey]; ok {
		command()
	} else {
		log.Fatalf("Error: Solution for day %d, part %d has not been implemented yet.", day, part)
	}
}
