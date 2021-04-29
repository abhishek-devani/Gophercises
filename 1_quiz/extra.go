package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func main() {

	filename := readArg()

	// fmt.Println(filename)

	file, err := openFile(filename)
	// fmt.Println("out" + fmt.Sprintln(file) + "\n")
	if err != nil {
		exit(fmt.Sprintf("Failed to oepn the CSV file: %s\n", filename))
	}

	lines, err := readCSV(file)

	problems, err := ParseLines(lines)

	score := printProblems(problems)

	total := len(problems)

	fmt.Printf("\nyou scored %d out of %d\n\n", score, total)

}

func readArg() string {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the formate of 'question, answer' ")
	flag.Parse()

	return *csvFileName
}

func openFile(filename string) (io.Reader, error) {
	// fmt.Println("\nin : " + fmt.Sprintln(os.Open(filename)) + "\n")
	return os.Open(filename)
}

func readCSV(file io.Reader) ([][]string, error) {
	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	// fmt.Println(lines)

	return lines, nil
}

func ParseLines(lines [][]string) ([]problem, error) {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	fmt.Println(ret)
	return ret, nil
}

func printProblems(problems []problem) int {

	count := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.answer {
			count++
		}
	}
	return count
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
