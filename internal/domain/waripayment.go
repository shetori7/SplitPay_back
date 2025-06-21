package domain

type Wari_payment struct {
	PaymentId    int     `gorm:"primaryKey;autoIncrement" json:"payment_id"`
	PayerGroupId string  `json:"group_uuid"`
	PayerUserId  int     `json:"user_id"`
	PayerAmount  float64 `json:"payer_amount"`
}
