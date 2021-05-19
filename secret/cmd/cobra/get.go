package cobra

import (
	"fmt"

	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/secret"
	"github.com/spf13/cobra"
)

var MockGet bool

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret from your secret storage.",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil || MockGet {
			fmt.Println("No Value Set")
			return
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
