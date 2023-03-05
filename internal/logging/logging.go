package logging

import (
	as "github.com/palavrapasse/aspirador/pkg"
)

var Aspirador as.Aspirador

func CreateAspiradorClients() []as.Client {
	consoleClient := as.NewConsoleClient()
	return []as.Client{&consoleClient}
}
