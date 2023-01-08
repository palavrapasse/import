package logging

import (
	as "github.com/palavrapasse/aspirador/pkg"
)

var Aspirador as.Aspirador

func CreateAspiradorClients() []as.Client {
	return []as.Client{as.NewConsoleClient()}
}
