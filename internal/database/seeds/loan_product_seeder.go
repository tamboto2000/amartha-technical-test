package seeds

import (
	"context"
	"database/sql"
)

func LoadProductSeeder(ctx context.Context, db *sql.DB) error {
	q := `INSERT INTO loan_products (name, description, amount, interest, installment, installment_period, tenor) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.ExecContext(ctx, q,
		"Pendanaan 555",
		"Dapatkan pinjaman 5JT Rupiah dengan tenor 50 minggu! Bunga hanya 10%!",
		5000000,
		10,
		110000,
		7,
		350,
	)

	return err
}
