package main

import (
	"os"
	"reflect"
	"testing"
)

func TestReadArg(t *testing.T) {
	inp, _ := ReadArg()
	exp := "problems.csv"
	if inp != exp {
		t.Fatal("error")
	}
}

func TestCode(t *testing.T) {

	path := "/home/gslab/Desktop/Gophercises/1_quiz/problems.csv"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatal("file is not exists")
	}

	f := "problems.csv"
	input, err := OpenFile(f)
	if err != nil {
		t.Fatal("error")
	}

	lines, err := ReadCSV(input)
	if err != nil {
		t.Fatal("error")
	}

	problems, err := ParseLines(lines)
	if err != nil {
		t.Fatal("error")
	}

	var tes []problem

	tes = append(tes, problem{question: "5+5", answer: "10"})
	tes = append(tes, problem{question: "1+1", answer: "2"})
	tes = append(tes, problem{question: "8+3", answer: "11"})
	tes = append(tes, problem{question: "1+2", answer: "3"})
	tes = append(tes, problem{question: "8+6", answer: "14"})
	tes = append(tes, problem{question: "3+1", answer: "4"})
	tes = append(tes, problem{question: "1+4", answer: "5"})
	tes = append(tes, problem{question: "5+1", answer: "6"})
	tes = append(tes, problem{question: "2+3", answer: "5"})
	tes = append(tes, problem{question: "3+3", answer: "6"})
	tes = append(tes, problem{question: "2+4", answer: "6"})
	tes = append(tes, problem{question: "5+2", answer: "7"})

	check := reflect.DeepEqual(problems, tes)

	if check != true {
		t.Fatal("does not match input and output")
	}

}
