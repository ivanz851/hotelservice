package booking

import (
	"database/sql"
	"fmt"
	"hotelservice/internal/models"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
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

func (s *Storage) GetBookings() ([]models.Booking, error) {
	rows, err := s.db.Query("SELECT id, hotel_id, client_id FROM bookings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(&booking.ID, &booking.HotelID, &booking.ClientID); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}

func (s *Storage) GetBooking(bookingID int) (*models.Booking, error) {
	row := s.db.QueryRow("SELECT id, hotel_id, client_id FROM bookings WHERE id = $1", bookingID)
	var booking models.Booking

	if err := row.Scan(&booking.ID, &booking.HotelID, &booking.ClientID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("booking with ID %d not found", bookingID)
		}
		return nil, err
	}
	return &booking, nil
}

func (s *Storage) AddBooking(booking models.Booking) error {
	_, err := s.db.Exec("INSERT INTO bookings (hotel_id, client_id, status, email) VALUES ($1, $2, $3, $4)", booking.HotelID, booking.ClientID, "confirmed", booking.Email)
	return err
}
