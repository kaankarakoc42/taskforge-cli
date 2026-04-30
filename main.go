package main

import (
	"log"

	"taskforge-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
