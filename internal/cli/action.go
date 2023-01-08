package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/palavrapasse/damn/pkg/entity"
	"github.com/palavrapasse/import/internal"
	"github.com/palavrapasse/import/internal/parser"
	"github.com/urfave/cli/v2"
)

func CreateAction(databasePath *string, leakPath *string, context *string, platforms *cli.StringSlice, shareDate *cli.Timestamp, leakers *cli.StringSlice, storeImport func(databasePath string, i entity.Import) error) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		internal.Aspirador.Warning("Starting Import")

		var errors []error

		err := validateFilePath(*leakPath, FlagLeakPath)
		errors = appendValidError(errors, err)

		err = validateFilePath(*databasePath, FlagDatabasePath)
		errors = appendValidError(errors, err)

		if len(errors) != 0 {
			return errors[0]
		}

		var parser parser.LeakParser = parser.PlainTextLeakParser{
			FilePath: *leakPath,
		}

		leakParse, errParse := parser.Parse()

		if errParse != nil {

			var errParseString string
			for _, v := range errParse {
				errParseString = fmt.Sprintf("%s\n%s", errParseString, v)
			}
			internal.Aspirador.Warning(fmt.Sprintf("Found the following errors parsing leak: %s", errParseString))

			fmt.Println("Proceed with import?")
			reader := bufio.NewReader(os.Stdin)
			input, _, errRead := reader.ReadLine()

			if errRead != nil {
				return errRead
			}

			if !IsProceedAnswer(proceedAnswers, strings.ToLower(string(input))) {
				internal.Aspirador.Info("Stopped import")
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

		leak, err := entity.NewLeak(*context, sharedatesc)
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

		err = storeImport(*databasePath, i)

		if err != nil {
			return err
		}

		internal.Aspirador.Info("Successful Import")
		return nil
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

func appendValidError(errors []error, err error) []error {

	if err != nil {
		errors = append(errors, err)
	}

	return errors
}
