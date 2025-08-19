package billing

import (
	"context"

	"github.com/tamboto2000/amartha-technical-test/internal/apps/billing/repositories"
)

type BillingService struct {
	billRepo repositories.BillRepository
}

func NewBillingService(billRepo repositories.BillRepository) BillingService {
	return BillingService{
		billRepo: billRepo,
	}
}

func (svc BillingService) IsDelinquent(ctx context.Context, userId int) (bool, error) {
	return svc.billRepo.IsUserDelinquent(ctx, userId)
}

func (svc BillingService) GetOutstanding(ctx context.Context, userId, userLoanId int) (int64, error) {
	return svc.billRepo.GetUserLoanOutstanding(ctx, userId, userLoanId)
}

func (svc BillingService) MakePayment(ctx context.Context, userId, userLoanId int, amount int64) error {
	bill, err := svc.billRepo.GetUserUpcomingUnpaidBill(ctx, userId, userLoanId)
	if err != nil {
		return err
	}

	repayment, err := bill.Pay(amount)
	if err != nil {
		return err
	}

	err = svc.billRepo.CreateRepayment(ctx, repayment)
	return err
}
