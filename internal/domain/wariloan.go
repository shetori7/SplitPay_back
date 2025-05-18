package domain

type Wari_loan struct {
	LoanId      string `gorm:"primaryKey;autoIncrement"`
	PaymentId   int
	PayeeUserId int
	PayeeAmount float64
}
