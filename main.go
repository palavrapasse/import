package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	// log.Println("** Import Project **")

	// var parser parser.LeakParser = parser.PlainTextLeakParser{
	// 	FilePath: "./plaintext.txt",
	// }

	// leakParse, errors := parser.Parse()
	// log.Println(errors)
	// log.Println(leakParse)

	app := &cli.App{
		Name:                 "import",
		Version:              "v0.0.1",
		Usage:                "Imports leak files into SQLite",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "leak",
				Aliases: []string{"l"},
				Usage:   "import leak from file",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() == 0 {
						fmt.Println("Command -> missing file")
						return nil
					}

					fmt.Println("Command -> import leak from file: ", cCtx.Args().First())
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load leak from `FILE`",
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.IsSet("config") {
				file := cCtx.String("config")
				fmt.Println("Flag -> Flag: ", file)
				fmt.Println("Flag -> import leak from file: ", file)
			}
			return nil
		},
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "current version",
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
