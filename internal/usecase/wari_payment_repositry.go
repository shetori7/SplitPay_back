package usecase

import (
	"SplitPay_back/internal/domain"
	"SplitPay_back/internal/dto"
)

type WariPaymentRepository interface {
	Store(*domain.Wari_payment)
	SelectByGroupId(groupId int) []domain.Wari_payment
	Select() []domain.Wari_payment
	Delete(id int) error
	SelectPaymentAndLoanByGroupId(groupUuid string) []dto.ReCalcFinalPaymentDto
}
