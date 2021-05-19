package cmd

import (
	"fmt"
	// "main/go/src/github.com/abhishek-devani/Gophercises/task/db"
	"strings"

	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/task/db"
	"github.com/spf13/cobra"
)

var MockAdd bool

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a tak to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		err := db.CreateTask(task)
		if err != nil || MockAdd {
			return
		}
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
