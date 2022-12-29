package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/palavrapasse/import/internal/database"
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

const (
	DBFilePath = "../fandom/database2.sqlite"
)

func main() {

	exampleCommand := fmt.Sprintf(`./import --leak-path="path/file.txt" --context="context" --platforms="platform1, platform2" --share-date="%s" --leakers="leaker1, leaker2"`,
		entity.DateFormatLayout)

	var leakPath string
	var context string
	var platforms cli.StringSlice
	var shareDate cli.Timestamp
	var leakers cli.StringSlice

	app := &cli.App{
		Name:                 "import",
		Version:              "v0.0.1",
		Usage:                "Imports leak files into SQLite",
		Copyright:            "(c) 2022 palavrapasse",
		Suggest:              true,
		EnableBashCompletion: true,
		HideHelp:             false,
		HideVersion:          false,
		Authors: []*cli.Author{
			{Name: "João Freitas"},
			{Name: "Rute Santos"},
		},
		Commands: []*cli.Command{},
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        FlagLeakPath,
				Aliases:     []string{"lp"},
				Usage:       "Load leak from `FILE`",
				Required:    true,
				Destination: &leakPath,
			},
			&cli.StringFlag{
				Name:        FlagLeakContext,
				Aliases:     []string{"c"},
				Usage:       "Leak Context",
				Required:    true,
				Destination: &context,
			},
			&cli.StringSliceFlag{
				Name:        FlagLeakPlatforms,
				Aliases:     []string{"p"},
				Usage:       "Platforms affected by the leak (comma separator)",
				Value:       cli.NewStringSlice("default"),
				Required:    false,
				Destination: &platforms,
			},
			&cli.TimestampFlag{
				Name:        FlagLeakShareDate,
				Aliases:     []string{"sd"},
				Usage:       "Leak Share Date",
				Layout:      entity.DateFormatLayout,
				Required:    true,
				Destination: &shareDate,
			},
			&cli.StringSliceFlag{
				Name:        FlagLeakers,
				Aliases:     []string{"l"},
				Usage:       "Leakers (comma separator)",
				Required:    true,
				Destination: &leakers,
			},
		},
		Action: func(cCtx *cli.Context) error {
			err := validateValue(leakPath, FlagLeakPath)

			if err != nil {
				return err
			}

			var parser parser.LeakParser = parser.PlainTextLeakParser{
				FilePath: leakPath,
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

			platformsSlice := platforms.Value()
			err = validateSlice(platformsSlice, FlagLeakPlatforms)

			if err != nil {
				return err
			}

			leakersSlice := leakers.Value()
			err = validateSlice(leakersSlice, FlagLeakers)

			if err != nil {
				return err
			}

			shareDateFormat := shareDate.Value().Format(entity.DateFormatLayout)

			sharedatesc, err := entity.NewDateInSeconds(shareDateFormat)

			if err != nil {
				return err
			}

			leak, err := entity.NewLeak(context, sharedatesc)

			if err != nil {
				return err
			}

			leakPlatforms, err := createPlatforms(platformsSlice)

			if err != nil {
				return err
			}

			leakBadActors, err := createBadActors(leakersSlice)

			if err != nil {
				return err
			}

			i := entity.Import{
				Leak:              leak,
				AffectedUsers:     leakParse,
				AffectedPlatforms: leakPlatforms,
				Leakers:           leakBadActors,
			}

			err = storeImport(i)

			if err != nil {
				return err
			}

			return nil
		},
	}

	// Append to an existing template
	cli.AppHelpTemplate = fmt.Sprintf(`%s
EXAMPLE: 
	%s

WEBSITE:
	https://github.com/palavrapasse

`, cli.AppHelpTemplate, exampleCommand)

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func storeImport(i entity.Import) error {
	dbctx, err := database.NewDatabaseContext(DBFilePath)

	if dbctx.DB != nil {
		defer dbctx.DB.Close()
	}

	if err != nil {
		return fmt.Errorf("error while creating DB context")
	}

	err = dbctx.Insert(i)

	if err != nil {
		return fmt.Errorf("error while storing data in DB %v", err)
	}

	log.Println("Successful Import")
	return nil
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

func validateSlice(value []string, flag string) error {
	if len(value) == 0 {
		return fmt.Errorf("%s should not be empty", flag)
	}

	return nil
}
