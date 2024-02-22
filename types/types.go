package types

type PriceResponse struct {
	Price  float64 `json:"price"`
	Ticker string  `json:"ticker"`
}
