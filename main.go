package main

import (
	"fmt"
	"os"

	as "github.com/palavrapasse/aspirador/pkg"
	"github.com/palavrapasse/damn/pkg/database"
	"github.com/palavrapasse/damn/pkg/entity"
	"github.com/palavrapasse/damn/pkg/entity/query"
	"github.com/palavrapasse/import/internal/cli"
	"github.com/palavrapasse/import/internal/http"
	"github.com/palavrapasse/import/internal/logging"
)

func main() {

	logging.Aspirador = as.WithClients(logging.CreateAspiradorClients())

	app := cli.CreateCliApp(storeImport, http.NotifyNewLeak)

	if err := app.Run(os.Args); err != nil {
		logging.Aspirador.Error(err.Error())
		return
	}
}

func storeImport(databasePath string, i query.Import) (entity.AutoGenKey, error) {
	logging.Aspirador.Info("Starting storage of Leak")

	dbctx, err := database.NewDatabaseContext[query.Import](databasePath)

	if dbctx.DB != nil {
		defer dbctx.DB.Close()
	}

	if err != nil {
		return entity.AutoGenKey(0), fmt.Errorf("could not open database connection: %w", err)
	}

	leakId, errInsert := dbctx.Insert(i)

	if errInsert != nil {
		return entity.AutoGenKey(0), fmt.Errorf("error while storing data in DB %w", errInsert)
	}

	logging.Aspirador.Info("Successful storage")

	return leakId, nil
}
