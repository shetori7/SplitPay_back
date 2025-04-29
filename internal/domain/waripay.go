package domain

import (
	"github.com/shopspring/decimal"
)

type Wari_payments struct {
	PaymentId string          `json:"payment_id"`
	GroupId   string          `json:"group_id"`
	UserId    string          `json:"user_id"`
	Amount    decimal.Decimal `json:"amount"`
}
