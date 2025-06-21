package database

import (
	"SplitPay_back/internal/domain"
)

type WariFinalPaymentRepository struct {
	SqlHandler
}

func (db *WariFinalPaymentRepository) StoreAll(wfp []*domain.WariFinalPayment) {
	db.Create(&wfp)
}
