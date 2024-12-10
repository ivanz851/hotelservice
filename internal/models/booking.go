package models

type Booking struct {
	ID                     int    `json:"id"`
	ClientID               int    `json:"client_id"`
	HotelID                int    `json:"hotel_id"`
	RoomID                 int    `json:"room_id"`
	BookingStartTimestamp  string `json:"booking_start_timestamp"`
	BookingFinishTimestamp string `json:"booking_finish_timestamp"`
	Date                   string `json:"creation_timestamp"`
}
