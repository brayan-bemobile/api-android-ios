package entities

type Transactions struct {
	Sku      string  `json:"sku,omitempty"`
	Amount   float64 `json:"amount,omitempty" form:"amount"`
	Currency string  `json:"currency,omitempty" form:"currency"`
}
