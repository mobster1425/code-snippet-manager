package cmd

import (
	"fmt"
	// "strings"

	"feyin/go-code-snippet-manager/db"
	"github.com/spf13/cobra"
)



var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a snippet from the code snippet manager.",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		category := args[1]

		code, err := db.GetSnippet(name, category)
		if err != nil {
			fmt.Println("Snippet not found:", err)
			return
		}

		fmt.Printf("Name: %s\nCategory: %s\nCode:\n%s\n", code.Name, code.Category, code.Code)
	},
}

// makes sure the add and get commands are initialized before the program is run
// so we would be able to use the commands while giving input
func init() {

	RootCmd.AddCommand(getCmd)
}