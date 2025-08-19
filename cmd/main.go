package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/tamboto2000/amartha-technical-test/internal/config"
	"github.com/tamboto2000/amartha-technical-test/internal/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	db, err := database.ConnectToDB(cfg.Database)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = database.RunMigration(context.Background(), cfg.Database, db)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = database.RunSeeder(context.Background(), db)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// Billing service
	// billsRepo := repositories.NewPQBillRepository(db)
	// billingSvc := billing.NewBillingService(billsRepo)

	// TODO: Use the IsDelinquent, GetOutstanding, and MakePayment
}
