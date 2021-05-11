package cmd

import (
	"fmt"
	"main/go/src/github.com/abhishek-devani/Gophercises/task/db"
	"strconv"

	"github.com/spf13/cobra"
)

var MockDo1 bool
var MockDo2 bool
var MockDo3 bool

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil || MockDo1 {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		var err error
		tasks, err := db.AllTasks()
		if err != nil || MockDo2 {
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid Task Number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTasks(task.Key)
			if err != nil || MockDo3 {
				return
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
