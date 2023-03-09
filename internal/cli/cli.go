package cli

import (
	"fmt"
	"sort"
	"time"

	"github.com/palavrapasse/damn/pkg/entity"
	"github.com/palavrapasse/damn/pkg/entity/query"
	"github.com/urfave/cli/v2"
)

func CreateCliApp(storeImport func(databasePath string, i query.Import) (entity.AutoGenKey, error), notifyImport func(entity.AutoGenKey, string) error) cli.App {

	var databasePath string
	var leakPath string
	var context string
	var platforms cli.StringSlice
	var shareDate cli.Timestamp
	var leakers cli.StringSlice
	var notifyNewLeakURL string
	var skipInteractiveMode bool

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
		Flags:                CreateCliFlags(&databasePath, &leakPath, &context, &platforms, &shareDate, &leakers, &notifyNewLeakURL, &skipInteractiveMode),
		Action:               CreateAction(&databasePath, &leakPath, &context, &platforms, &shareDate, &leakers, &notifyNewLeakURL, &skipInteractiveMode, storeImport, notifyImport),
	}

	cli.AppHelpTemplate = CreateAppHelpTemplate(cli.AppHelpTemplate)

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return *app
}
