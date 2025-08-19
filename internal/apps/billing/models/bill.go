package models

import (
	"errors"
	"time"
)

type Bill struct {
	ID      int
	Number  string
	Amount  int64
	DueDate time.Time
}

type Repayment struct {
	ID     int
	BillID int
	Amount int64
}

func (b Bill) Pay(amount int64) (Repayment, error) {
	if amount != b.Amount {
		return Repayment{}, errors.New("payment amount did not match with bill amount")
	}

	return Repayment{
		BillID: b.ID,
		Amount: amount,
	}, nil
}
