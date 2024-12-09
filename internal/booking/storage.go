package booking

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

type Booking struct {
	ID       int    `json:"id"`
	HotelID  int    `json:"hotel_id"`
	ClientID int    `json:"client_id"`
	Date     string `json:"date"`
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
	rows, err := s.db.Query("SELECT bookingId, hotelId, clientId, creationTimestamp FROM Bookings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var booking Booking
		if err := rows.Scan(&booking.ID, &booking.HotelID, &booking.ClientID, &booking.Date); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}

func (s *Storage) AddBooking(booking Booking) error {
	_, err := s.db.Exec("INSERT INTO Bookings (hotelId, clientId, creationTimestamp) VALUES ($1, $2, $3)", booking.HotelID, booking.ClientID, booking.Date)
	return err
}
