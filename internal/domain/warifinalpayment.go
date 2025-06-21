package domain

type WariFinalPayment struct {
	FinalPaymentId int     `gorm:"primaryKey;autoIncrement" json:"final_payment_id"`
	GroupUuid      string  `json:"group_uuid"`
	FromUserId     int     `json:"from_user_id"`
	ToUserId       int     `json:"to_user_id"`
	Amount         float64 `json:"amount"`
}
