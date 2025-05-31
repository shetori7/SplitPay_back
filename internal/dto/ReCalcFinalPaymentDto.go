package dto

type ReCalcFinalPaymentDto struct {
	PaymentId    int     `gorm:"primaryKey;autoIncrement" json:"payment_id"`
	PayerGroupId int     `json:"payer_group_id"`
	PayerUserId  int     `json:"payer_user_id"`
	PayerAmount  float64 `json:"amount"`
	Payment_date string  `json:"payment_date"`
	Message      string  `json:"message"`
	Loan_id      int     `json:"loan_id"`
	PayeeUserId  int     `json:"payee_user_id"`
	PayeeAmount  float64 `json:"payee_amount"`
}
