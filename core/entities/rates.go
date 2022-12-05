package entities

type Rates struct {
	From string  `json:"from,omitempty"`
	To   string  `json:"to,omitempty" form:"to"`
	Rate float64 `json:"rate,omitempty" form:"rate"`
}
