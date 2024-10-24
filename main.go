package main

import (
	"log"

	"proxx-test-task/pkg/ui"
)

func main() {
	if err := ui.RunCliView(); err != nil {
		log.Fatal(err)
	}
}
