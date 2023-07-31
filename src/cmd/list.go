package cmd


import (
	"fmt"
	// "strings"

	"feyin/go-code-snippet-manager/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the snippets in the code snippet manager.",
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := db.ListSnippets()
		if err != nil {
			fmt.Println("Error listing snippets:", err)
			return
		}

		if len(snippets) == 0 {
			fmt.Println("No snippets found.")
			return
		}

		fmt.Println("=== Snippets ===")
		for _, s := range snippets {
			fmt.Printf("Name: %s\nCategory: %s\nCode:\n%s\n\n", s.Name, s.Category, s.Code)
		}
	},
}


// makes sure the add, get, and list commands are initialized before the program is run
// so we would be able to use the commands while giving input
func init() {
	RootCmd.AddCommand(listCmd)
}