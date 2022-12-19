package main

import (
	"log"

	"github.com/palavrapasse/import/internal"
)

func main() {
	log.Println("** Import Project **")

	baId := internal.AutoGenKey(10)
	log.Println(baId)

	badActor := internal.BadActor{
		BaId:       baId,
		Identifier: "Identifier",
	}
	log.Println(badActor)
}
