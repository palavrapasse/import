package main

import (
	"log"

	"github.com/palavrapasse/import/internal/parser"
)

func main() {
	log.Println("** Import Project **")

	var parser parser.LeakParser = parser.PlainTextLeakParser{
		FilePath: "./plaintext.txt",
	}

	leakParse, errors := parser.Parse()
	log.Println(errors)
	log.Println(leakParse)
}
