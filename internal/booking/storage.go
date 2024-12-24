package booking

import (
	"database/sql"
	_ "github.com/lib/pq"
	"hotelservice/internal/models"
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
	rows, err := s.db.Query("SELECT id, hotel_id, client_id FROM Bookings")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

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

func (s *Storage) AddBooking(booking models.Booking) error {
	_, err := s.db.Exec("INSERT INTO bookings (hotel_id, client_id, room_id) VALUES ($1, $2, $3)", booking.HotelID, booking.ClientID)
	return err
}
