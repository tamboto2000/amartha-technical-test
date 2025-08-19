package database

import (
	"context"
	"database/sql"

	"github.com/tamboto2000/amartha-technical-test/internal/database/seeds"
)

// TODO: Using goose to run the seeder
func RunSeeder(ctx context.Context, db *sql.DB) error {
	if err := seeds.UserSeeder(ctx, db); err != nil {
		return err
	}

	if err := seeds.LoadProductSeeder(ctx, db); err != nil {
		return err
	}

	return nil
}
