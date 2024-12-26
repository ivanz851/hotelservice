package models

type Booking struct {
	ID       int    `json:"id"`
	ClientID int    `json:"client_id"`
	HotelID  int    `json:"hotel_id"`
	Status   string `json:"status"`
	Email    string `json:"email"`
}
