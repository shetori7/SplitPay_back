package usecase

import (
	"SplitPay_back/internal/domain"
)

type PaymentRepository interface {
	Store(*domain.Wari_payments)
	Select() []domain.Wari_payments
	Delete(id string)
}
