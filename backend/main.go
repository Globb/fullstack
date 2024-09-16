package main

import (
	"os"

	"fullstack/backend/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
