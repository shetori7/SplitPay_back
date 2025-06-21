package request

type PaymentNewRequestBody struct {
	GroupUuid      string  `json:"group_uuid"`
	PayerId        int     `json:"payer_id"`
	Amount         float64 `json:"amount"`
	ParticipantIds []int   `json:"participants_ids"`
}
