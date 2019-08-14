package main

import (
	"fmt"
	"os"
	"path/filepath"
	"stark/gophercises/task/cmd"
	"stark/gophercises/task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "task.db")
	must(db.Init(dbPath))
	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
