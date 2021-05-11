package main

import (
	"testing"
)

// func TestDo(t *testing.T) {
// 	cmd := exec.Command("./main", "do", "3")

// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = cmd.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	buffer := bufio.NewReader(stdout)
// 	var str string

// 	for {
// 		line, _, err := buffer.ReadLine()

// 		str = str + string(line)
// 		// fmt.Println(string(line))

// 		if err == io.EOF {
// 			break
// 		}
// 	}

// 	exp := "Marked \"3\" as completed."
// 	if str != exp {
// 		t.Fatalf("%v\n%v\n", str, exp)
// 	}
// }

// func TestList(t *testing.T) {
// 	cmd := exec.Command("./main", "list")

// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = cmd.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	buffer := bufio.NewReader(stdout)
// 	var str string

// 	for {
// 		line, _, err := buffer.ReadLine()

// 		str = str + string(line)
// 		// fmt.Println(string(line))

// 		if err == io.EOF {
// 			break
// 		}
// 	}

// 	exp := "You have the following tasks1. hahahaha2. hahaha added"
// 	if str != exp {
// 		t.Fatalf("%v\n%v\n", str, exp)
// 	}
// }

func TestMain(t *testing.T) {

	main()
	Mock = true
	main()
}
