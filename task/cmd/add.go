package cmd

import (
	"fmt"
	"main/go/src/github.com/abhishek-devani/Gophercises/task/db"
	"strings"

	"github.com/spf13/cobra"
)

var task string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a tak to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task = strings.Join(args, " ")
		err := db.CreateTask(task)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

func check() {
	RootCmd.AddCommand(addCmd)
}

func init() {
	check()
}
