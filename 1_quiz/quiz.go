package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {

	filename, timeL := ReadArg()

	file, err := OpenFile(filename)
	if err != nil {
		Exit(fmt.Sprintf("Failed to oepn the CSV file: %s\n", filename))
	}

	lines, _ := ReadCSV(file)

	problems, _ := ParseLines(lines)

	score := PrintProblems(problems, timeL)

	total := len(problems)

	fmt.Printf("\nyou scored %d out of %d\n\n", score, total)

}

func ReadArg() (string, int) {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the formate of 'question, answer' ")

	timeLimit := flag.Int("limit", 30, "time in seconds")

	flag.Parse()

	return *csvFileName, *timeLimit
}

func OpenFile(filename string) (io.Reader, error) {
	return os.Open(filename)
}

func ReadCSV(file io.Reader) ([][]string, error) {
	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		Exit("Failed to parse the provided CSV file")
	}

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
	return ret, nil
}

func PrintProblems(problems []problem, timeL int) int {

	timer := time.NewTimer(time.Duration(timeL) * time.Second)

	count := 0

	for i, p := range problems {

		fmt.Printf("Problem #%d: %s = ", i+1, p.question)

		answerCh := make(chan string)

		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()

		select {
		case <-timer.C:
			return count
		case answer := <-answerCh:
			if answer == p.answer {
				count++
			}
		}
	}
	return count
}

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
