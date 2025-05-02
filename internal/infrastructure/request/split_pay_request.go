package request

type Payment struct {
	Payer   string   `json:"payer"`
	Amount  float64  `json:"amount"`
	Members []string `json:"Members"`
}
