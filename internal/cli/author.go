package cli

import "github.com/urfave/cli/v2"

func CreateCliAuthors() []*cli.Author {

	return []*cli.Author{
		{Name: "João Freitas"},
		{Name: "Rute Santos"},
	}
}
