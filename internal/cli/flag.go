package cli

import (
	"github.com/palavrapasse/damn/pkg/entity/query"
	"github.com/urfave/cli/v2"
)

const (
	FlagDatabasePath     = "database-path"
	FlagLeakPath         = "leak-path"
	FlagLeakContext      = "context"
	FlagLeakPlatforms    = "platforms"
	FlagLeakShareDate    = "share-date"
	FlagLeakers          = "leakers"
	FlagNotifyNewLeakURL = "notify url"
)

func CreateCliFlags(databasePath *string, leakPath *string, context *string, platforms *cli.StringSlice, shareDate *cli.Timestamp, leakers *cli.StringSlice, notifyNewLeakURL *string) []cli.Flag {

	return []cli.Flag{
		&cli.PathFlag{
			Name:        FlagDatabasePath,
			Aliases:     AliasesFlagDatabasePath,
			Usage:       "Store leaks into `SQLite Database`",
			Required:    true,
			Destination: databasePath,
		},
		&cli.PathFlag{
			Name:        FlagLeakPath,
			Aliases:     AliasesFlagLeakPath,
			Usage:       "Load leak from `FILE`",
			Required:    true,
			Destination: leakPath,
		},
		&cli.StringFlag{
			Name:        FlagLeakContext,
			Aliases:     AliasesFlagLeakContext,
			Usage:       "Leak Context",
			Required:    true,
			Destination: context,
		},
		&cli.StringSliceFlag{
			Name:        FlagLeakPlatforms,
			Aliases:     AliasesFlagLeakPlatforms,
			Usage:       "Platforms affected by the leak (separated by commas)",
			Value:       cli.NewStringSlice("Unknown"),
			Required:    false,
			Destination: platforms,
		},
		&cli.TimestampFlag{
			Name:        FlagLeakShareDate,
			Aliases:     AliasesFlagLeakShareDate,
			Usage:       "Leak Share Date",
			Layout:      query.DateFormatLayout,
			Required:    true,
			Destination: shareDate,
		},
		&cli.StringSliceFlag{
			Name:        FlagLeakers,
			Aliases:     AliasesFlagLeakers,
			Usage:       "Leakers (separated by commas)",
			Required:    true,
			Destination: leakers,
		},
		&cli.StringFlag{
			Name:        FlagNotifyNewLeakURL,
			Aliases:     AliasesFlagNotifyNewLeakURL,
			Usage:       "URL service to be notified of the new leak",
			Required:    true,
			Destination: notifyNewLeakURL,
		},
	}
}
