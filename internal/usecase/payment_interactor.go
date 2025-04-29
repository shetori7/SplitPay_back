package usecase

import (
	"SplitPay_back/internal/domain"

	"github.com/google/uuid"
)

type PaymentInteractor struct {
	PaymentRepository PaymentRepository
}

func (interactor *PaymentInteractor) Add(g *domain.Wari_payments) {
	newUUID := uuid.New().String()
	g.PaymentId = newUUID
	interactor.PaymentRepository.Store(g)
}

func (interactor *PaymentInteractor) CalucuratePayment(amount float64, members int) float64 {
	if members == 0 {
		return 0
	}
	return amount / float64(members)
