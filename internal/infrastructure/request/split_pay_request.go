package request

type Payment struct {
	Payer        string   `json:"payer"`
	Amount       float64  `json:"amount"`
	Participants []string `json:"participants"`
}
