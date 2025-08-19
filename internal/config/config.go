package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Env keys for database config
const (
	envKeyDbUsername   = "AMARTHA_DATABASE_USERNAME"
	envKeyDbPass       = "AMARTHA_DATABASE_PASSWORD"
	envKeyDbHost       = "AMARTHA_DATABASE_HOST"
	envKeyDbPort       = "AMARTHA_DATABASE_PORT"
	envKeyDbDatabase   = "AMARTHA_DATABASE_DATABASE"
	envKeyDbSslMode    = "AMARTHA_DATABASE_SSL_MODE"
	envKeyMigrationDir = "AMARTHA_DATABASE_MIGRATION_DIR"
)

// Env keys for app config
const (
	envKeyAppHttpPort = "AMARTHA_APP_HTTP_PORT"
	envKeyAppLogLevel = "AMARTHA_APP_LOG_LEVEL"
)

// Default values for app config
const (
	appCfgDefHttpPort = 8080
	appCfgDefLogLevel = slog.LevelInfo
)

type Database struct {
	Username     string
	Password     string
	Host         string
	Port         string
	Database     string
	SSLMode      string
	MigrationDir string
}

type App struct {
	HTTPPort int
	LogLevel slog.Level
}

type Config struct {
	Database Database
	App      App
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}, fmt.Errorf("error on loading config: %v", err)
	}

	var cfg Config
	dbCfg := loadDatabaseCfg()
	appCfg, err := loadAppCfg()
	if err != nil {
		return cfg, err
	}

	cfg.Database = dbCfg
	cfg.App = appCfg

	return cfg, nil
}

func loadDatabaseCfg() Database {
	return Database{
		Username:     os.Getenv(envKeyDbUsername),
		Password:     os.Getenv(envKeyDbPass),
		Host:         os.Getenv(envKeyDbHost),
		Port:         os.Getenv(envKeyDbPort),
		Database:     os.Getenv(envKeyDbDatabase),
		SSLMode:      os.Getenv(envKeyDbSslMode),
		MigrationDir: os.Getenv(envKeyMigrationDir),
	}
}

func loadAppCfg() (App, error) {
	httpPort, err := parseIntEnv(envKeyAppHttpPort, appCfgDefHttpPort)
	if err != nil {
		return App{}, err
	}

	// TODO: Maybe try to use string and map it to slog log levels
	logLevelInt, err := parseIntEnv(envKeyAppLogLevel, int(appCfgDefLogLevel))
	if err != nil {
		return App{}, err
	}

	return App{
		HTTPPort: httpPort,
		LogLevel: slog.Level(logLevelInt),
	}, nil
}

func parseIntEnv(key string, def int) (int, error) {
	val := os.Getenv(key)
	if val == "" {
		return def, nil
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("error on parsing env %s: %v", key, err)
	}

	return i, nil
}
