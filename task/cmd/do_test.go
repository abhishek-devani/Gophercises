package cmd

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"testing"
)

func TestDo(t *testing.T) {

	db := startDB()
	defer db.Close()

	a := []string{"1"}
	doCmd.Run(doCmd, a)

}

func TestDoOutput(t *testing.T) {
	cmd := exec.Command("./main", "do", "3")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	buffer := bufio.NewReader(stdout)
	var str string

	for {
		line, _, err := buffer.ReadLine()

		str = str + string(line)
		// fmt.Println(string(line))

		if err == io.EOF {
			break
		}
	}

	exp := "Marked \"3\" as completed."
	if str != exp {
		t.Fatalf("%v\n%v\n", str, exp)
	}
}

func TestMock(t *testing.T) {

	db := startDB()
	defer db.Close()

	a := []string{"1"}
	MockDo1 = true
	doCmd.Run(doCmd, a)
	MockDo1 = false

	MockDo2 = true
	doCmd.Run(doCmd, a)
	MockDo2 = false

	MockDo3 = true
	doCmd.Run(doCmd, a)
	MockDo3 = false

	b := []string{"10000000"}
	doCmd.Run(doCmd, b)
}
