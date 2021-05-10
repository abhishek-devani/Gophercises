package main

import (
	"main/go/src/github.com/abhishek-devani/Gophercises/task/cmd"
	"main/go/src/github.com/abhishek-devani/Gophercises/task/db"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}
