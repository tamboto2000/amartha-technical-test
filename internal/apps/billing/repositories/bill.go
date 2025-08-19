package repositories

import (
	"context"
	"database/sql"

	"github.com/tamboto2000/amartha-technical-test/internal/apps/billing/models"
	"github.com/tamboto2000/amartha-technical-test/internal/customerror"
)

type BillRepository interface {
	IsUserDelinquent(ctx context.Context, userId int) (bool, error)
	GetUserLoanOutstanding(ctx context.Context, userId, userLoanId int) (int64, error)
	GetUserUpcomingUnpaidBill(ctx context.Context, userId, userLoanId int) (models.Bill, error)
	CreateRepayment(ctx context.Context, repayment models.Repayment) error
}

type PQBillRepository struct {
	db *sql.DB
}

func NewPQBillRepository(db *sql.DB) PQBillRepository {
	return PQBillRepository{db: db}
}

func (repo PQBillRepository) IsUserDelinquent(ctx context.Context, userId int) (bool, error) {
	q := `SELECT     
    (COUNT(*) >= 2) AS is_delinquent
FROM bills b
LEFT JOIN repayments r ON b.id = r.bill_id
WHERE 
	b.user_id = $1
	AND b.due_date < CURRENT_DATE
  	AND r.id IS NULL
GROUP BY b.user_id;`

	row := repo.db.QueryRowContext(ctx, q, userId)
	var isDelinquent bool
	if err := row.Scan(&isDelinquent); err != nil {
		if err == sql.ErrNoRows {
			return false, customerror.ErrNotFound
		}

		return false, err
	}

	return isDelinquent, nil
}

func (repo PQBillRepository) GetUserLoanOutstanding(ctx context.Context, userId, userLoanId int) (int64, error) {
	q := `SELECT 
    COALESCE(SUM(b.amount), 0) AS total_outstanding
FROM bills b
LEFT JOIN repayments r ON b.id = r.bill_id
WHERE r.id IS NULL
  	AND b.user_id = $1
  	AND b.user_loan_id = $2`

	row := repo.db.QueryRowContext(ctx, q, userId, userLoanId)
	var outstanding int64
	if err := row.Scan(&outstanding); err != nil {
		if err == sql.ErrNoRows {
			return 0, customerror.ErrNotFound
		}

		return 0, err
	}

	return outstanding, nil
}

func (repo PQBillRepository) GetUserUpcomingUnpaidBill(ctx context.Context, userId, userLoanId int) (models.Bill, error) {
	q := `SELECT 
	b.id,
	b.number,
	b.amount,
	b.due_date
FROM bills b
LEFT JOIN repayments r ON b.id = r.bill_id
WHERE b.user_id = $1
  AND b.user_loan_id = $2
  AND r.id IS NULL          -- not yet paid
ORDER BY b.due_date ASC
LIMIT 1;`

	row := repo.db.QueryRowContext(ctx, q, userId, userLoanId)
	var bill models.Bill
	if err := row.Scan(
		&bill.ID,
		&bill.Number,
		&bill.Amount,
		&bill.DueDate,
	); err != nil {
		if err == sql.ErrNoRows {
			return models.Bill{}, customerror.ErrNotFound
		}

		return models.Bill{}, err
	}

	return bill, nil
}

func (repo PQBillRepository) CreateRepayment(ctx context.Context, repayment models.Repayment) error {
	q := `INSERT INTO repayments (bill_id, amount) VALUES ($1, $2)`
	_, err := repo.db.ExecContext(ctx, q, repayment.BillID, repayment.Amount)

	return err
}
