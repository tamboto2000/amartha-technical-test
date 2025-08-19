package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/tamboto2000/amartha-technical-test/internal/config"
)

func ConnectToDB(cfg config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.SSLMode,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("error on opening connection to database: %v", err)
	}

	return db, nil
}
