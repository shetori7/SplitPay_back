package request

type PaymentNewRequestBody struct {
	GroupId        int     `json:"group_id"`
	PayerId        int     `json:"payer_id"`
	Amount         float64 `json:"amount"`
	ParticipantIds []int   `json:"participants_ids"`
}
