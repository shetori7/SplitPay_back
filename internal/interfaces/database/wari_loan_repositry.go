package database

import "SplitPay_back/internal/domain"

type WariLoanRepository struct {
	SqlHandler
}

func (db *WariLoanRepository) Store(wp *domain.Wari_loan) {
	db.Create(&wp)
}
