package seeds

import (
	"context"
	"database/sql"

	"github.com/tamboto2000/amartha-technical-test/internal/helper/password"
)

type user struct {
	name             string
	email            string
	password         []byte
	passwordSalt     []byte
	credibilityLevel int
}

func UserSeeder(ctx context.Context, db *sql.DB) error {
	// Just use the same password, this is just for demonstration purposes
	pass := "12345678"
	passHash, passSalt, err := password.HashPassword(pass, 32, password.DefaultCost)
	if err != nil {
		return err
	}

	// Insert a dummy user complete with password
	users := []user{
		{
			name:             "Rizal",
			email:            "rizal@gmail.com",
			password:         passHash,
			passwordSalt:     passSalt,
			credibilityLevel: 0,
		},
		{
			name:             "Fajri",
			email:            "fajri@gmail.com",
			password:         passHash,
			passwordSalt:     passSalt,
			credibilityLevel: 0,
		},
	}

	q := `INSERT INTO users (name, email, password, password_salt, credibility_level) VALUES ($1, $2, $3, $4, $5)`

	for _, u := range users {
		_, err := db.ExecContext(ctx, q, u.name, u.email, u.password, u.passwordSalt, u.credibilityLevel)
		if err != nil {
			return err
		}
	}

	return nil
}
