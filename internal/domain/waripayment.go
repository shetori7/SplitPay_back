package domain

type Wari_payment struct {
	PaymentId    int     `gorm:"primaryKey;autoIncrement" json:"payment_id"`
	PayerGroupId int     `json:"group_id"`
	PayerUserId  int     `json:"user_id"`
	PayerAmount  float64 `json:"payer_amount"`
}
