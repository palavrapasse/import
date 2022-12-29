package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/palavrapasse/import/internal/database"
	"github.com/palavrapasse/import/internal/entity"
	"github.com/palavrapasse/import/internal/parser"
	"github.com/urfave/cli/v2"
)

const (
	FlagDatabasePath  = "database-path"
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

var ProceedAnswers = []string{ProceedShortAnswer, ProceedLongAnswer}
var ExampleCommand = fmt.Sprintf(`./import --database-path=""path/db.sqlite" --leak-path="path/file.txt" --context="context" --platforms="platform1, platform2" --share-date="%s" --leakers="leaker1, leaker2"`,
	entity.DateFormatLayout)

func main() {

	var databasePath string
	var leakPath string
	var context string
	var platforms cli.StringSlice
	var shareDate cli.Timestamp
	var leakers cli.StringSlice

	app := &cli.App{
		Name:                 "import",
		Version:              "v0.0.1",
		Usage:                "Imports leak files into SQLite",
		Copyright:            fmt.Sprintf("(c) %d palavrapasse", time.Now().Year()),
		Suggest:              true,
		EnableBashCompletion: true,
		HideHelp:             false,
		HideVersion:          false,
		Authors:              createCliAuthors(),
		Commands:             []*cli.Command{},
		Flags:                createCliFlags(&databasePath, &leakPath, &context, &platforms, &shareDate, &leakers),
		Action: func(cCtx *cli.Context) error {

			var errors []error

			err := validateFilePath(leakPath, FlagLeakPath)
			errors = appendValidError(errors, err)

			err = validateFilePath(databasePath, FlagDatabasePath)
			errors = appendValidError(errors, err)

			if len(errors) != 0 {
				return errors[0]
			}

			var parser parser.LeakParser = parser.PlainTextLeakParser{
				FilePath: leakPath,
			}

			leakParse, errParse := parser.Parse()

			if errParse != nil {
				log.Println("Found the following errors parsing leak:")

				for _, v := range errParse {
					log.Println(v)
				}

				log.Println("Proceed with import?")
				reader := bufio.NewReader(os.Stdin)
				input, _, errRead := reader.ReadLine()

				if errRead != nil {
					return errRead
				}

				if !contains(ProceedAnswers, strings.ToLower(string(input))) {
					return nil
				}
			}

			platformsSlice := platforms.Value()
			err = validateFlagValues(platformsSlice, FlagLeakPlatforms)
			errors = appendValidError(errors, err)

			leakersSlice := leakers.Value()
			err = validateFlagValues(leakersSlice, FlagLeakers)
			errors = appendValidError(errors, err)

			leakPlatforms, err := createPlatforms(platformsSlice)
			errors = appendValidError(errors, err)

			leakBadActors, err := createBadActors(leakersSlice)
			errors = appendValidError(errors, err)

			shareDateFormat := shareDate.Value().Format(entity.DateFormatLayout)
			sharedatesc, err := entity.NewDateInSeconds(shareDateFormat)
			errors = appendValidError(errors, err)

			if len(errors) != 0 {
				return errors[0]
			}

			leak, err := entity.NewLeak(context, sharedatesc)
			errors = appendValidError(errors, err)

			if len(errors) != 0 {
				return errors[0]
			}

			i := entity.Import{
				Leak:              leak,
				AffectedUsers:     leakParse,
				AffectedPlatforms: leakPlatforms,
				Leakers:           leakBadActors,
			}

			err = storeImport(databasePath, i)

			if err != nil {
				return err
			}

			log.Println("Successful Import")
			return nil
		},
	}

	cli.AppHelpTemplate = createAppHelpTemplate(cli.AppHelpTemplate)

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func createCliAuthors() []*cli.Author {

	return []*cli.Author{
		{Name: "João Freitas"},
		{Name: "Rute Santos"},
	}
}

func createCliFlags(databasePath *string, leakPath *string, context *string, platforms *cli.StringSlice, shareDate *cli.Timestamp, leakers *cli.StringSlice) []cli.Flag {

	return []cli.Flag{
		&cli.PathFlag{
			Name:        FlagDatabasePath,
			Aliases:     []string{"db"},
			Usage:       "Store leaks into `SQLite Database`",
			Required:    true,
			Destination: databasePath,
		},
		&cli.PathFlag{
			Name:        FlagLeakPath,
			Aliases:     []string{"lp"},
			Usage:       "Load leak from `FILE`",
			Required:    true,
			Destination: leakPath,
		},
		&cli.StringFlag{
			Name:        FlagLeakContext,
			Aliases:     []string{"c"},
			Usage:       "Leak Context",
			Required:    true,
			Destination: context,
		},
		&cli.StringSliceFlag{
			Name:        FlagLeakPlatforms,
			Aliases:     []string{"p"},
			Usage:       "Platforms affected by the leak (separated by commas)",
			Value:       cli.NewStringSlice("default"),
			Required:    false,
			Destination: platforms,
		},
		&cli.TimestampFlag{
			Name:        FlagLeakShareDate,
			Aliases:     []string{"sd"},
			Usage:       "Leak Share Date",
			Layout:      entity.DateFormatLayout,
			Required:    true,
			Destination: shareDate,
		},
		&cli.StringSliceFlag{
			Name:        FlagLeakers,
			Aliases:     []string{"l"},
			Usage:       "Leakers (separated by commas)",
			Required:    true,
			Destination: leakers,
		},
	}
}

func createAppHelpTemplate(base string) string {

	// Append to an existing template
	return fmt.Sprintf(`%s
EXAMPLE: 
	%s

WEBSITE:
	https://github.com/palavrapasse

`, base, ExampleCommand)
}

func storeImport(databasePath string, i entity.Import) error {
	dbctx, err := database.NewDatabaseContext(databasePath)

	if dbctx.DB != nil {
		defer dbctx.DB.Close()
	}

	if err != nil {
		return fmt.Errorf("could not open database connection: %v", err)
	}

	err = dbctx.Insert(i)

	if err != nil {
		return fmt.Errorf("error while storing data in DB %v", err)
	}

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

func validateFilePath(value string, flag string) error {
	if len(strings.TrimSpace(value)) == 0 {
		return fmt.Errorf("%s should not be empty or only white spaces", flag)
	}

	return nil
}

func validateFlagValues(value []string, flag string) error {
	if len(value) == 0 {
		return fmt.Errorf("%s should not be empty", flag)
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {

		if a == e {
			return true
		}
	}

	return false
}

func appendValidError(errors []error, err error) []error {

	if err != nil {
		errors = append(errors, err)
	}

	return errors
}
