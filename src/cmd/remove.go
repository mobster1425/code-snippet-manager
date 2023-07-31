package cmd


import (
	"fmt"
	// "strings"

	"feyin/go-code-snippet-manager/db"
	"github.com/spf13/cobra"
)


var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a snippet from the code snippet manager.",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		category := args[1]

		key, err := db.RemoveSnippet(name, category)
		if err != nil {
			fmt.Println("Snippet not found:", err)
			return
		}

		fmt.Printf("Removed \"%s\" from the \"%s\" category with key \"%s\".\n", name, category, key)
	},
}


// makes sure the add, get, list, and remove commands are initialized before the program is run
// so we would be able to use the commands while giving input
func init() {

	RootCmd.AddCommand(removeCmd)
}