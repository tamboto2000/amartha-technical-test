package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/tamboto2000/amartha-technical-test/internal/config"
)

func RunMigration(ctx context.Context, cfg config.Database, db *sql.DB) error {
	err := goose.UpContext(ctx, db, cfg.MigrationDir)
	if err != nil {
		return fmt.Errorf("error on running database migration: %v", err)
	}

	return nil
}
