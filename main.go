package main

import (
	"log"
	"os"
	"sort"
	"strings"
	"time"

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
	var platforms []string
	var sharedate string
	var leakers []string

	app := &cli.App{
		Name:                 "import",
		Version:              "v0.0.1",
		Usage:                "Imports leak files into SQLite",
		EnableBashCompletion: true,
		Commands:             []*cli.Command{},
		Flags: []cli.Flag{
			&cli.PathFlag{
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
				Name:     "context",
				Aliases:  []string{"c"},
				Usage:    "Leak Context",
				Required: true,
				Action: func(ctx *cli.Context, v string) error {
					validateValue(v, "context")
					context = v
					return nil
				},
			},
			&cli.StringSliceFlag{
				Name:     "platforms",
				Aliases:  []string{"p"},
				Usage:    "Platforms affected by the leak (comma separator)",
				Required: true,
				Action: func(ctx *cli.Context, v []string) error {
					validateValues(v, "platforms")
					platforms = v
					return nil
				},
			},
			&cli.TimestampFlag{
				Name:     "share-date",
				Aliases:  []string{"sd"},
				Usage:    "Leak Share Date",
				Layout:   entity.DateFormatLayout,
				Required: true,
				Action: func(ctx *cli.Context, v *time.Time) error {
					sharedate = v.Format(entity.DateFormatLayout)
					return nil
				},
			},
			&cli.StringSliceFlag{
				Name:     "leakers",
				Aliases:  []string{"l"},
				Usage:    "Leakers (comma separator)",
				Required: true,
				Action: func(ctx *cli.Context, v []string) error {
					validateValues(v, "leakers")
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

	if err := app.Run(os.Args); err != nil {
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
		log.Printf("sharedatesc -> %v\n", sharedatesc)

		leak, err := entity.NewLeak(context, sharedatesc)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(leak)

		createPlatforms(platforms)

		createBadActors(leakers)
	}
}

func createPlatforms(platforms []string) []entity.Platform {
	var list []entity.Platform

	for _, v := range platforms {
		platform, err := entity.NewPlatform(v)

		if err != nil {
			log.Fatal(err)
		}

		list = append(list, platform)
		log.Println(platform)
	}

	return list
}

func createBadActors(leakers []string) []entity.BadActor {
	var list []entity.BadActor

	for _, v := range leakers {
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

func validateValues(value []string, flag string) {
	if len(value) == 0 {
		log.Fatal(flag + " should not be empty")
	}
}
