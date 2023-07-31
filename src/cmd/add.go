package cmd

import (
	"fmt"
	"strings"

	"feyin/go-code-snippet-manager/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a snippet to the code snippet manager.",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		category := args[1]
		code := strings.Join(args[2:], " ")

		key, err := db.CreateSnippet(name, code, category)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		fmt.Printf("Added \"%s\" to the \"%s\" category with key \"%s\".\n", name, category, key)
	},
}

// makes sure the add command is initialized before the program is run
// so we would be able to use the add command while giving input
func init() {
	RootCmd.AddCommand(addCmd)
}
