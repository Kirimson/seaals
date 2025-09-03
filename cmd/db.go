package cmd

import (
	"context"
	"seaals/config"
	"seaals/models"

	"github.com/urfave/cli/v3"
)

func Migrate(ctx context.Context, cmd *cli.Command) error {
	database := cmd.String("database")

	sealDB, err := models.InitialiseDB(database)
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate schemas
	if _, err := sealDB.ExecContext(ctx, config.Ddl); err != nil {
		return err
	}

	return nil
}
