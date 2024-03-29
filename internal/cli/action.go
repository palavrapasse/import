package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/palavrapasse/damn/pkg/entity"
	"github.com/palavrapasse/damn/pkg/entity/query"
	"github.com/palavrapasse/import/internal/logging"
	"github.com/palavrapasse/import/internal/parser"
	"github.com/urfave/cli/v2"
)

const MaxErrorLogCalls = 20000

func CreateAction(databasePath *string, leakPath *string, context *string, platforms *cli.StringSlice,
	shareDate *cli.Timestamp, leakers *cli.StringSlice, notifyNewLeakURL *string, skipInteractiveMode *bool,
	storeImport func(databasePath string, i query.Import) (entity.AutoGenKey, error),
	notifyImport func(entity.AutoGenKey, string) error,
) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		logging.Aspirador.Info("Starting Import")

		var errors []error

		err := validateNonEmptyValue(*leakPath, FlagLeakPath)
		errors = appendValidError(errors, err)

		err = validateNonEmptyValue(*databasePath, FlagDatabasePath)
		errors = appendValidError(errors, err)

		err = validateNonEmptyValue(*notifyNewLeakURL, FlagNotifyNewLeakURL)
		errors = appendValidError(errors, err)

		if len(errors) != 0 {
			return errors[0]
		}

		var parser parser.LeakParser = parser.PlainTextLeakParser{
			FilePath: *leakPath,
		}

		leakParse, errParse := parser.Parse()

		if errParse != nil {
			errorsCount := len(errParse)

			if errorsCount > MaxErrorLogCalls {
				logging.Aspirador.Warning(fmt.Sprintf("Found a lot of errors during leak parse (%d)...", errorsCount))
			} else {
				logging.Aspirador.Warning("Found the following errors parsing leak:")

				for _, v := range errParse {
					logging.Aspirador.Warning(v.Error())
				}
			}

			if !*skipInteractiveMode {
				fmt.Println("Proceed with import?")
				reader := bufio.NewReader(os.Stdin)
				input, _, errRead := reader.ReadLine()

				if errRead != nil {
					return errRead
				}

				if !IsProceedAnswer(proceedAnswers, strings.ToLower(string(input))) {
					logging.Aspirador.Info("Stopped import")
					return nil
				}
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

		shareDateFormat := shareDate.Value().Format(query.DateFormatLayout)
		sharedatesc, err := query.NewDateInSeconds(shareDateFormat)
		errors = appendValidError(errors, err)

		if len(errors) != 0 {
			return errors[0]
		}

		leak, err := query.NewLeak(*context, sharedatesc)
		errors = appendValidError(errors, err)

		if len(errors) != 0 {
			return errors[0]
		}

		i := query.Import{
			Leak:              leak,
			AffectedUsers:     leakParse,
			AffectedPlatforms: leakPlatforms,
			Leakers:           leakBadActors,
		}

		leakId, errImport := storeImport(*databasePath, i)

		if errImport != nil {
			return errImport
		}

		logging.Aspirador.Info(fmt.Sprintf("Successful Import (%d)", len(leakParse)))

		err = notifyImport(leakId, *notifyNewLeakURL)

		if err != nil {
			return err
		}

		return nil
	}
}

func createPlatforms(platforms []string) ([]query.Platform, error) {
	var list []query.Platform

	for _, v := range platforms {
		platform, err := query.NewPlatform(v)

		if err != nil {
			return list, err
		}

		list = append(list, platform)
	}

	return list, nil
}

func createBadActors(leakers []string) ([]query.BadActor, error) {
	var list []query.BadActor

	for _, v := range leakers {
		badActor, err := query.NewBadActor(v)

		if err != nil {
			return list, err
		}

		list = append(list, badActor)
	}

	return list, nil
}

func validateNonEmptyValue(value string, flag string) error {
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
