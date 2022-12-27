package main

import (
	"log"
	"os"
	"sort"
	"strings"

	"github.com/palavrapasse/import/internal/entity"
	"github.com/palavrapasse/import/internal/parser"
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

	var path string
	var context string
	var platforms string
	var sharedate string
	var leakers string

	app := &cli.App{
		Name:                 "import",
		Version:              "v0.0.1",
		Usage:                "Imports leak files into SQLite",
		EnableBashCompletion: true,
		Commands:             []*cli.Command{},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "leak-path",
				Aliases:  []string{"lp"},
				Usage:    "Load leak from `FILE`",
				Required: true,
				Action: func(ctx *cli.Context, v string) error {
					validateValue(v, "leak-path")
					path = v
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "context",
				Aliases: []string{"c"},
				Usage:   "Context",
				Action: func(ctx *cli.Context, v string) error {
					validateValue(v, "context")
					context = v
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "platforms",
				Aliases: []string{"p"},
				Usage:   "Platforms",
				Action: func(ctx *cli.Context, v string) error {
					validateValue(v, "platforms")
					platforms = v
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "share-date",
				Aliases: []string{"sd"},
				Usage:   "Share Date",
				Action: func(ctx *cli.Context, v string) error {
					validateValue(v, "share-date")
					sharedate = v
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "leakers",
				Aliases: []string{"l"},
				Usage:   "Leakers",
				Action: func(ctx *cli.Context, v string) error {
					validateValue(v, "leakers")
					leakers = v
					return nil
				},
			},
		},
		Action: func(cCtx *cli.Context) error {
			//Without this action the help message will be shown
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if path != "" {
		log.Println(path)
		log.Println(context)
		log.Println(platforms)
		log.Println(sharedate)
		log.Println(leakers)

		var parser parser.LeakParser = parser.PlainTextLeakParser{
			FilePath: path,
		}
		leakParse, errors := parser.Parse()
		log.Println(errors)
		log.Println(leakParse)

		sharedatesc, err := entity.NewDateInSeconds(sharedate)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(sharedatesc)

		leak, err := entity.NewLeak(context, sharedatesc)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(leak)

		createPlatforms(platforms)

		createBadActors(leakers)
	}
}

func createPlatforms(platforms string) []entity.Platform {
	platformsSplit := strings.Split(platforms, ",")
	var list []entity.Platform

	for _, v := range platformsSplit {
		platform, err := entity.NewPlatform(v)

		if err != nil {
			log.Fatal(err)
		}

		list = append(list, platform)
		log.Println(platform)
	}

	return list
}

func createBadActors(leakers string) []entity.BadActor {
	leakersSplit := strings.Split(leakers, ",")
	var list []entity.BadActor

	for _, v := range leakersSplit {
		badActor, err := entity.NewBadActor(v)

		if err != nil {
			log.Fatal(err)
		}

		list = append(list, badActor)
		log.Println(badActor)
	}

	return list
}

func validateValue(value string, flag string) {
	if strings.TrimSpace(value) == "" {
		log.Fatal(flag + " should not be empty or white spaces")
	}
}
