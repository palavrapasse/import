package main

import (
	"fmt"
	"log"
	"os"

	"github.com/palavrapasse/damn/pkg/database"
	"github.com/palavrapasse/damn/pkg/entity"
	"github.com/palavrapasse/import/internal/cli"
)

func main() {

	app := cli.CreateCliApp(storeImport)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
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
