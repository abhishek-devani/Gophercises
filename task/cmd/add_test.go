package cmd

import (
	"main/go/src/github.com/abhishek-devani/Gophercises/task/db"
	"path/filepath"
	"testing"

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

	MockAdd = true

	addCmd.Run(addCmd, a)
}
