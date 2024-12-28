package storages

import "time"

type ExchangeRate struct {
	ID           int       `json:"id"`
	FromCurrency string    `json:"from_currency"`
	ToCurrency   string    `json:"to_currency"`
	Rate         float64   `json:"rate"`
	UpdatedAt    time.Time `json:"autoUpdateTime"`
}
