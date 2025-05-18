package usecase

import (
	"SplitPay_back/internal/domain"
)

type WariLoanRepository interface {
	Store(*domain.Wari_loan)
}
