package models

type Booking struct {
	ID                     int    `json:"bookingId"`
	ClientID               int    `json:"clientId"`
	HotelID                int    `json:"hotelId"`
	RoomID                 int    `json:"roomId"`
	RoomCategory           string `json:"roomCategory"`
	Price                  int    `json:"price"`
	BookingStartTimestamp  string `json:"bookingStartTimestamp"`
	BookingFinishTimestamp string `json:"bookingFinishTimestamp"`
	Date                   string `json:"creationTimestamp"`
}
