package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Code Snippet Manager",
	Short: "A Simple Cli Manager to Manage Your Codes Snippets",
}