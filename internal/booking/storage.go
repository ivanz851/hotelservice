package booking

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

type Booking struct {
	ID                     int    `json:"id"`
	ClientID               int    `json:"client_id"`
	HotelID                int    `json:"hotel_id"`
	RoomID                 int    `json:"room_id"`
	BookingStartTimestamp  string `json:"booking_start_timestamp"`
	BookingFinishTimestamp string `json:"booking_finish_timestamp"`
	Date                   string `json:"creation_timestamp"`
}

func NewStorage(conn string) *Storage {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic("Connection Failed: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("Unable to connect to the database: " + err.Error())
	}

	return &Storage{db: db}
}

func (s *Storage) GetBookings() ([]Booking, error) {
	rows, err := s.db.Query("SELECT id, hotel_id, client_id, room_id, creation_timestamp FROM Bookings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var booking Booking
		if err := rows.Scan(&booking.ID, &booking.HotelID, &booking.ClientID, &booking.RoomID, &booking.Date); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}

func (s *Storage) AddBooking(booking Booking) error {
	_, err := s.db.Exec("INSERT INTO bookings (hotel_id, client_id, room_id, creation_timestamp) VALUES ($1, $2, $3, $4)", booking.HotelID, booking.ClientID, booking.Date)
	return err
}
