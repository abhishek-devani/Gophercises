package cmd

import (
	"fmt"
	// "main/go/src/github.com/abhishek-devani/Gophercises/task/db"

	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/task/db"
	"github.com/spf13/cobra"
)

var MockList1 bool
var MockList2 bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil || MockList1 {
			fmt.Println(err)
			return
		}
		if len(tasks) == 0 || MockList2 {
			fmt.Println("You have no task to complete! yeah")
			return
		}
		fmt.Println("You have the following tasks")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
