package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/palavrapasse/import/internal/entity"
	"github.com/palavrapasse/import/internal/parser"
	"github.com/urfave/cli/v2"
)

const (
	FlagLeakPath      = "leak-path"
	FlagLeakContext   = "context"
	FlagLeakPlatforms = "platforms"
	FlagLeakShareDate = "share-date"
	FlagLeakers       = "leakers"
)

const (
	ProceedShortAnswer = "y"
	ProceedLongAnswer  = "yes"
)

func main() {

	app := &cli.App{
		Name:                 "import",
		Version:              "v0.0.1",
		Usage:                "Imports leak files into SQLite",
		Copyright:            "(c) 2022 palavrapasse",
		Suggest:              true,
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{Name: "João Freitas"},
			{Name: "Rute Santos"},
		},
		Commands: []*cli.Command{},
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     FlagLeakPath,
				Aliases:  []string{"lp"},
				Usage:    "Load leak from `FILE`",
				Required: true,
			},
			&cli.StringFlag{
				Name:     FlagLeakContext,
				Aliases:  []string{"c"},
				Usage:    "Leak Context",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     FlagLeakPlatforms,
				Aliases:  []string{"p"},
				Usage:    "Platforms affected by the leak (comma separator)",
				Value:    cli.NewStringSlice("default"),
				Required: false,
			},
			&cli.TimestampFlag{
				Name:     FlagLeakShareDate,
				Aliases:  []string{"sd"},
				Usage:    "Leak Share Date",
				Layout:   entity.DateFormatLayout,
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     FlagLeakers,
				Aliases:  []string{"l"},
				Usage:    "Leakers (comma separator)",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			path := cCtx.String(FlagLeakPath)
			err := validateValue(path, FlagLeakPath)

			if err != nil {
				return err
			}

			var parser parser.LeakParser = parser.PlainTextLeakParser{
				FilePath: path,
			}

			leakParse, errors := parser.Parse()

			if errors != nil {
				log.Println("Found the following errors in the file:")

				for _, v := range errors {
					log.Println(v)
				}

				log.Println("Should the import be proceeded?")
				reader := bufio.NewReader(os.Stdin)
				input, _, errRead := reader.ReadLine()

				if errRead != nil {
					return errRead
				}

				if string(input) != ProceedShortAnswer && string(input) != ProceedLongAnswer {
					return nil
				}
			}
			log.Println(errors)
			log.Println(leakParse)

			platforms := cCtx.StringSlice(FlagLeakPlatforms)
			err = validateValues(platforms, FlagLeakPlatforms)

			if err != nil {
				return err
			}

			leakers := cCtx.StringSlice(FlagLeakers)
			err = validateValues(leakers, FlagLeakers)

			if err != nil {
				return err
			}

			context := cCtx.String(FlagLeakContext)
			sharedate := cCtx.Timestamp(FlagLeakShareDate).Format(entity.DateFormatLayout)

			sharedatesc, err := entity.NewDateInSeconds(sharedate)

			if err != nil {
				return err
			}

			leak, err := entity.NewLeak(context, sharedatesc)

			if err != nil {
				return err
			}

			leakPlatforms, err := createPlatforms(platforms)

			if err != nil {
				return err
			}

			leakBadActors, err := createBadActors(leakers)

			if err != nil {
				return err
			}

			log.Println(leak)
			log.Println(leakPlatforms)
			log.Println(leakBadActors)

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

func createPlatforms(platforms []string) ([]entity.Platform, error) {
	var list []entity.Platform

	for _, v := range platforms {
		platform, err := entity.NewPlatform(v)

		if err != nil {
			return list, err
		}

		list = append(list, platform)
	}

	return list, nil
}

func createBadActors(leakers []string) ([]entity.BadActor, error) {
	var list []entity.BadActor

	for _, v := range leakers {
		badActor, err := entity.NewBadActor(v)

		if err != nil {
			return list, err
		}

		list = append(list, badActor)
	}

	return list, nil
}

func validateValue(value string, flag string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s should not be empty or white spaces", flag)
	}

	return nil
}

func validateValues(value []string, flag string) error {
	if len(value) == 0 {
		return fmt.Errorf("%s should not be empty", flag)
	}

	return nil
}
