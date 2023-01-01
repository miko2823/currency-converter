package models

type LatestRates struct {
	Base   string `json:"base"`
	Amount int    `json:"amount"`
	Rates  []Rate `json:"rates"`
}

type Rate struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}
