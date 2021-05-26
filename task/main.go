package main

import (

	// "main/go/src/github.com/abhishek-devani/Gophercises/task/cmd"
	// "main/go/src/github.com/abhishek-devani/Gophercises/task/db"

	"path/filepath"

	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/task/cmd"
	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/task/db"
	"github.com/mitchellh/go-homedir"
)

var Mock bool

func main() {
	home, _ := homedir.Dir()
	// fmt.Println(home)
	dbPath := filepath.Join(home, "tasks.db")
	_, err := db.Init(dbPath)
	if err != nil || Mock {
		return
	}
	cmd.RootCmd.Execute()
}
