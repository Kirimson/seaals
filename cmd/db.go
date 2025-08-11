package cmd

import (
	"context"
	"fmt"
	"seaals/models"
	"time"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func migrateTable(db *gorm.DB, m any, t string) error {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf("Migrating %s table", t)
	s.FinalMSG = fmt.Sprintf("Finished migrating %s table\n", t)
	s.Start()
	defer s.Stop()
	if err := db.AutoMigrate(&m, &models.Tag{}); err != nil {
		s.FinalMSG = fmt.Sprintf("Failed to migrate %s table\n", t)
		return err
	}
	return nil
}

func Migrate(ctx context.Context, cmd *cli.Command) error {
	database := cmd.String("database")
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate schemas
	if err := migrateTable(db, &models.Seal{}, "seals"); err != nil {
		return err
	}
	if err := migrateTable(db, &models.Tag{}, "tags"); err != nil {
		return err
	}
	return nil
}
