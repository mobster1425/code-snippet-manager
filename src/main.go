package main

import (
	"fmt"
	"os"
	"path/filepath"

	"feyin/go-code-snippet-manager/cmd"
	"feyin/go-code-snippet-manager/db"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	// getting the home directory on the user system
	home, _ := homedir.Dir()
	// joining the homedirectory with code-snippets.db
	dbPath := filepath.Join(home, "code-snippets.db")

	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

// the must function is used for handling error
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
