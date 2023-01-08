package internal

import (
	as "github.com/palavrapasse/aspirador/pkg"
)

var Aspirador as.Aspirador

func init() {
	Aspirador = as.NewAspirador()
}
