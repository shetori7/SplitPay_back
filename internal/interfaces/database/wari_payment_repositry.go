package database

import (
	"SplitPay_back/internal/domain"
	"SplitPay_back/internal/dto"
)

type WariPaymentRepository struct {
	SqlHandler
}

func (db *WariPaymentRepository) Store(wp *domain.Wari_payment) {
	db.Create(&wp)
}

func (db *WariPaymentRepository) SelectByGroupId(groupId int) []domain.Wari_payment {
	payments := []domain.Wari_payment{}
	db.FindById(&payments, groupId)
	return payments
}

func (db *WariPaymentRepository) SelectPaymentAndLoanByGroupId(groupUuid string) []dto.ReCalcFinalPaymentDto {
	reCalcFinalPaymentDtos := []dto.ReCalcFinalPaymentDto{}
	db.Raw().Table("wari_payments").
		Select("wari_payments.*, wari_loans.*").
		Joins("JOIN wari_loans ON wari_payments.payment_id = wari_loans.payment_id").
		Where("wari_payments.payer_group_id = ?", groupUuid).
		Scan(&reCalcFinalPaymentDtos)
	return reCalcFinalPaymentDtos
}

func (db *WariPaymentRepository) Select() []domain.Wari_payment {
	payments := []domain.Wari_payment{}
	db.FindAll(&payments)
	return payments
}

func (db *WariPaymentRepository) Delete(id int) error {
	payments := []domain.Wari_payment{}
	return db.DeleteById(&payments, id)
}

func (db *WariPaymentRepository) DeleteByUuid(groupUuid string) error {
	payments := []domain.Wari_payment{}
	result := db.Raw().Table("wari_final_payments").
		Where("group_uuid = ?", groupUuid).Delete(&payments)
	return result.Error
}
