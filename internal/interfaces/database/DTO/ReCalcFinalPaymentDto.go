package dto

type ReCalcFinalPaymentDto struct {
	PaymentId    int     `gorm:"primaryKey;autoIncrement" json:"payment_id"`
	PayerGroupId int     `json:"payer_group_id"`
	PayerUserId  int     `json:"payer_user_id"`
	Amount       float64 `json:"amount"`
	Payment_date string  `json:"payment_date"`
	Message      string  `json:"message"`
	Loan_id      int     `json:"loan_id"`
	ToUserId     int     `json:"to_user_id"`
}
