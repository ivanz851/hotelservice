package models

type Hotel struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Address       string  `json:"address"`
	PricePerNight float64 `json:"price_per_night"`
}
