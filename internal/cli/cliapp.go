package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/palavrapasse/import/internal/entity"
	"github.com/palavrapasse/import/internal/parser"
	"github.com/urfave/cli/v2"
)

const (
	proceedShortAnswer = "y"
	proceedLongAnswer  = "yes"
)

var proceedAnswers = []string{proceedShortAnswer, proceedLongAnswer}
var exampleCommand = fmt.Sprintf(`./import --database-path="path/db.sqlite" --leak-path="path/file.txt" --context="context" --platforms="platform1, platform2" --share-date="%s" --leakers="leaker1, leaker2"`,
	entity.DateFormatLayout)

func CreateCliApp(storeImport func(databasePath string, i entity.Import) error) cli.App {

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
		Authors:              CreateCliAuthors(),
		Commands:             []*cli.Command{},
		Flags:                CreateCliFlags(&databasePath, &leakPath, &context, &platforms, &shareDate, &leakers),
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

				fmt.Println("Proceed with import?")
				reader := bufio.NewReader(os.Stdin)
				input, _, errRead := reader.ReadLine()

				if errRead != nil {
					return errRead
				}

				if !contains(proceedAnswers, strings.ToLower(string(input))) {
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

	return *app
}

func createAppHelpTemplate(base string) string {

	// Append to an existing template
	return fmt.Sprintf(`%s
EXAMPLE: 
	%s

WEBSITE:
	https://github.com/palavrapasse

`, base, exampleCommand)
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
