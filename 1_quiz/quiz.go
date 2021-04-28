package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// flag created
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the formate of 'question, answer' ")
	flag.Parse()

	openCSV(*csvFileName)

}

func openCSV(csvFileName string) {

	file, err := os.Open(csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to oepn the CSV file: %s\n", csvFileName))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	problems := ParseLines(lines)

	printProblems(problems, lines)

}

func printProblems(problems []problem, lines [][]string) {
	count := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.answer {
			count++
		}
	}
	fmt.Printf("\nYou scored %d out of %d\n\n", count, len(lines))
}

func ParseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
