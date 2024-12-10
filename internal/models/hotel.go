package models

type Hotel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Hotelier int    `json:"hotelier"`
	Rating   int    `json:"rating"`
	Country  string `json:"country"`
	Address  string `json:"address"`
}
