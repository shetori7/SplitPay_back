package usecase

import (
	"SplitPay_back/internal/domain"
)

type WariFinalPaymentRepository interface {
	StoreAll([]*domain.WariFinalPayment)
}
