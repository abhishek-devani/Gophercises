package cobra

import (
	"fmt"

	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/secret"
	"github.com/spf13/cobra"
)

var MockSet bool

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage.",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil || MockSet {
			return
		}
		fmt.Printf("Value Set Successfully!\n")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
