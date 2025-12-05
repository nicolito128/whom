package main

import (
	"log"

	"github.com/nicolito128/whom/internal/gens"
)

func main() {
	err := gens.GenerateBaseProject("test_project")
	if err != nil {
		log.Fatal(err)
	}
}
