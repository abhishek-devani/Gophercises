package cmd

import (
	"bufio"
	"io"
	"log"

	// "main/go/src/github.com/abhishek-devani/Gophercises/task/db"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/task/db"
	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
)

func startDB() *bolt.DB {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "tasks.db")
	db, _ := db.Init(DbPath)
	return db
}

func TestAdd(t *testing.T) {
	db := startDB()
	defer db.Close()

	a := []string{"Watch Golang tutorial"}
	addCmd.Run(addCmd, a)

}

func TestAddOutput(t *testing.T) {
	cmd := exec.Command("./main", "add", "hahaha")

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

	exp := "Added \"hahaha\" to your task list"
	if str != exp {
		t.Fatalf("%v\n%v\n", str, exp)
	}
}

func TestMockAdd(t *testing.T) {
	db := startDB()
	defer db.Close()

	a := []string{"Watch Golang tutorial"}
	MockAdd = true
	addCmd.Run(addCmd, a)
}
