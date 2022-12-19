package main

import (
	"log"

	"github.com/palavrapasse/import/internal/entity"
)

func main() {
	log.Println("** Import Project **")

	baId := entity.AutoGenKey(10)
	log.Println(baId)

	badActor := entity.BadActor{
		BaId:       baId,
		Identifier: "Identifier",
	}
	log.Println(badActor)
}
