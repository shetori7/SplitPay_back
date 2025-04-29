package database

import (
	"SplitPay_back/internal/domain"
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

func (db *WariPaymentRepository) Select() []domain.Wari_payment {
	payments := []domain.Wari_payment{}
	db.FindAll(&payments)
	return payments
}

func (db *WariPaymentRepository) Delete(id string) {
	payments := []domain.Wari_payment{}
	db.DeleteById(&payments, id)
}
